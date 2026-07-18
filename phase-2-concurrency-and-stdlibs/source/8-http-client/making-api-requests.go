package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
	Problem: Here, we are using http.Get function, this function works fine but, but it will not close request in case the
	server took too long to respond.

	Sol: use http.Client{timeout}
*/

func BasicApiGetRequest() {
	// Basic get request, res will hold our connection, err in case if our connection fails
	res, err := http.Get("http://localhost:8080/dev-akshat")

	// If err, no connection is created simply exit.
	// If err is nil, it means connection is established so we must close it
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
	defer res.Body.Close()

	// Always check the StatusCode of API request to ensure we only process successful responses
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body) // body => slice of raw bytes

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}

		fmt.Println(string(body)) // converting raw slices of data into strings
	}
}

func BasicApiGetRequestWithTimeout() {
	client := http.Client{Timeout: 1 * time.Second}
	res, err := client.Get("http://localhost:8080/dev-akshat")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(-1)
		}

		fmt.Println(string(body))
	}
}
