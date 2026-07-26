package downloader

import (
	"context"
	"fmt"
	"net/http"
)

func simpleDownloadingRequest(ctx context.Context, client *http.Client, fileUrl string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("simpleDownloadingRequest: error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "concurrent-file-downloader/1.0")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("simpleDownloadingRequest: error executing request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("simpleDownloadingRequest: request failed with status code %d: %w", res.StatusCode, err)
	}

	return res, nil
}

func getHeadersFromServer(ctx context.Context, client *http.Client, fileUrl string) (http.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("getHeadersFromServer: error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "concurrent-file-downloader/1.0")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getHeadersFromServer: error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusPartialContent {
		return nil, fmt.Errorf("getHeadersFromServer: request failed with status code %d: %d", res.StatusCode, err)
	}

	return res.Header, nil
}

func downloadSingleChunkRequest(ctx context.Context, client *http.Client, fileUrl string, firstChunk int64, lastChunk int64) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("downloadSingleChunkRequest: error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "concurrent-file-downloader/1.0")
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", firstChunk, lastChunk))

	resp, err := client.Do(req);
	if err != nil {
		return nil, fmt.Errorf("downloadSingleChunkRequest: error executing request: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		return nil, fmt.Errorf("downloadSingleChunkRequest: request failed with status code %d: %w", resp.StatusCode, err)
	}

	return resp, nil
}
