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
)

var destPath = "./"

func SimpleFileDownloader(ctx context.Context, client *http.Client, srcUrl string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	res, err := simpleDownloadingRequest(ctx, client, srcUrl)
	if err != nil {
		return fmt.Errorf("SimpleFileDownloader-1: %w", err)
	}
	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	parsedURL, _ := url.Parse(srcUrl)
	fileName := path.Base(parsedURL.Path)

	expectedSize, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		// header missing or malformed — can't validate size, not necessarily a fatal error
		expectedSize = -1 // or 0, some sentinel meaning "unknown"
	}

	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	if err == nil {
		fileName = params["filename"]
	}

	f, err := os.Create(destPath + fileName)
	if err != nil {
		return fmt.Errorf("SimpleFileDownloader-2: %w", err)
	}
	defer f.Close()

	written, err := io.Copy(f, res.Body)

	if expectedSize > 0 && written != expectedSize {
		return fmt.Errorf("downloadFile: incomplete download, got %d of %d bytes", written, expectedSize)
	}

	fmt.Printf("Downloaded file %s of %fmb", fileName, float64(written)/1024.0/1024.0)
	return nil
}
