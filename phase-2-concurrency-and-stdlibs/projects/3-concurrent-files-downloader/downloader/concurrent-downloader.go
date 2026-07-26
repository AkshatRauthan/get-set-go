package downloader

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type fileStatus int

const (
	DownloadingSuccessful fileStatus = iota + 1 // Download completed
	DownloadPending                             // The download is under process
	PartialFailure                              // Some chunks failed
	RetryingDownload                            // Retrying download
	CompleteFailure                             // All chunks failed
)

type chunk struct {
	Idx        int   `json:"chunkIdx"`
	Start      int64 `json:"start"`
	End        int64 `json:"end"`
	Downloaded bool  `json:"downloaded"`
}

type progressFileStruct struct {
	Filename  string `json:"fileName"`
	FileUrl   string `json:"fileUrl"`
	Size      int64  `json:"size"`
	NumChunks int    `json:"numChunks"`

	Status          fileStatus `json:"status"`
	CompletedSize   int64      `json:"completedSize"`
	CompletedChunks int64      `json:"completedChunks"`

	Chunks []chunk `json:"chunks"`
}

func createNewProgressReport(filename string, fileUrl string, size int64, numChunks int) *progressFileStruct {
	return &progressFileStruct{
		Filename:  filename,
		FileUrl:   fileUrl,
		Size:      size,
		NumChunks: numChunks,
		Chunks:    make([]chunk, numChunks),
		Status:    DownloadPending,
	}
}

func (p *progressFileStruct) save() error {
	ext := filepath.Ext(p.Filename)
	baseName := strings.TrimSuffix(p.Filename, ext)
	outputPath := filepath.Join(".", baseName+".alan")

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("save: error while marshaling progress report: %w", err)
	}

	if err = os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("save: error while saving file: %w", err)
	}

	return nil
}

func loadProgressReport(filePath string) (*progressFileStruct, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("loadProgressReport: error while reading file: %w", err)
	}

	var report progressFileStruct
	if err = json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("loadProgressReport: error while unmarshaling file contents: %w", err)
	}

	return &report, nil
}

func ChunkedFileDownloader(ctx context.Context, client *http.Client, fileUrl string, chunkSize int64, maxConnections int) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	headers, err := getHeadersFromServer(ctx, client, fileUrl)
	if err != nil {
		return fmt.Errorf("ChunkedFileDownload-1: %w", err)
	}

	parsedURL, _ := url.Parse(fileUrl)
	fileName := path.Base(parsedURL.Path)

	contentLen, err := strconv.ParseInt(headers.Get("Content-Length"), 10, 64)
	if err != nil {
		return fmt.Errorf("ChunkedFileDownload: error while parsing content length value: %w", err)
	}

	if acceptRange := headers.Get("Accept-Ranges"); acceptRange != "bytes" {
		return fmt.Errorf("ChunkedFileDownload: accept-ranges header not supported on url")
	}

	_, params, err := mime.ParseMediaType(headers.Get("Content-Disposition"))
	if err == nil {
		fileName = params["filename"]
	}

	numOfChunks := int(math.Ceil(float64(contentLen) / float64(chunkSize)))

	connectionCount := make(chan struct{}, maxConnections)
	defer close(connectionCount)

	// download chunks and assemble them
	wg := sync.WaitGroup{}
	wg.Add(numOfChunks)

	destPath := path.Join("./" + fileName)
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("SimpleFileDownloader-2: %w", err)
	}
	defer destFile.Close()

	var mu sync.Mutex
	var chunkErrors []error

	var progressReport = createNewProgressReport(destPath, fileUrl, contentLen, numOfChunks)
	defer func() {
		err = progressReport.save()
		if err != nil {
			fmt.Println("Error during saving progress report")
		} else {
			fmt.Println("Progress report saved successfully")
		}
	}()

	start, end := int64(0), chunkSize-1
	for i := 1; i <= numOfChunks; i++ {
		if i == numOfChunks {
			end = contentLen - 1
		}

		go func(start int64, end int64, chunkIdx int) {
			defer wg.Done()

			connectionCount <- struct{}{}
			defer func() { <-connectionCount }()

			progressReport.Chunks[chunkIdx-1] = chunk{
				chunkIdx - 1,
				start,
				end,
				false,
			}

			resp, err := downloadSingleChunkRequest(ctx, client, fileUrl, start, end)
			if err != nil {
				mu.Lock()
				defer mu.Unlock()
				chunkErrors = append(chunkErrors, fmt.Errorf("ChunkedFileDownload-1: chunkId %d: %w", chunkIdx, err))

				return
			}
			defer resp.Body.Close()

			buf := make([]byte, 32*1024)
			offset := start

			for {
				n, readErr := resp.Body.Read(buf)
				if n > 0 {
					if _, writeErr := destFile.WriteAt(buf[:n], offset); writeErr != nil {
						mu.Lock()
						chunkErrors = append(chunkErrors, fmt.Errorf("ChunkedFileDownload-2: write error in chunk %d: %w", chunkIdx, writeErr))
						mu.Unlock()
						return
					}
					offset += int64(n)
				}
				if readErr == io.EOF {
					break
				}
				if readErr != nil {
					fmt.Printf("ChunkedFileDownload: read error in chunk %d: %v\n", chunkIdx, readErr)
					return
				}
			}

			progressReport.Chunks[chunkIdx-1].Downloaded = true
			fmt.Printf("Completed chunk %d: Written %.3f MBs\n", chunkIdx, float32(end-start)/1024.0/1024.0)

		}(start, end, i)

		start += chunkSize
		end += chunkSize
	}
	wg.Wait()

	if len(chunkErrors) > 0 {
		progressReport.Status = PartialFailure
		return fmt.Errorf("ChunkedFileDownload: %d chunk(s) failed: %#v", len(chunkErrors), chunkErrors)
	}

	fmt.Println("\nDone")

	return nil
}
