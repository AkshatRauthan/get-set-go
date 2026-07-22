package downloader

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
)

func ChunkedFileDownload(ctx context.Context, client *http.Client, fileUrl string, numOfChunks int) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	headers, err := getHeadersFromServer(ctx, client, fileUrl)
	if err != nil {
		return fmt.Errorf("ChunkedFileDownload-1: %w", err)
	}

	fmt.Print("\nHeaders:\n")

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

	chunkSize := contentLen / int64(numOfChunks)

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

	start, end := int64(0), chunkSize-1
	for i := 1; i <= numOfChunks; i++ {
		if i == numOfChunks {
			end = contentLen - 1
		}

		go func(start int64, end int64, chunkIdx int) {
			defer wg.Done()
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

			fmt.Printf("Completed chunk %d: Written %d MBs\n", chunkIdx, (end-start)/1024/1024)

		}(start, end, i)

		start += chunkSize
		end += chunkSize
	}
	wg.Wait()

	if len(chunkErrors) > 0 {
		return fmt.Errorf("ChunkedFileDownload: %d chunk(s) failed: %#v", len(chunkErrors), chunkErrors)
	}

	fmt.Println("\nDone")

	return nil
}
