package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func CustomGetRequestWithQuery() {
	url := "https://jsonplaceholder.typicode.com/comments"
	client := http.Client{Timeout: time.Second * 2}

	// For GET requests body is nil, so We don't defer call req.Body.Close()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("postId", "1")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing response body: %v\n", err)
		return
	}

	fmt.Printf("\nResponse Body:\n%s\n", string(body))
}
