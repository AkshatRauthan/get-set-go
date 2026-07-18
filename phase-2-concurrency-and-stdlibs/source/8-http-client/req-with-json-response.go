package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Go only serialize fields that start with Capital, so each field name should be uppercase.
type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func ApiReqWithJsonResponse() {
	const url = "https://jsonplaceholder.typicode.com"

	client := http.Client{Timeout: 1 * time.Second}
	res, err := client.Get(url + "/todos/4")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {

		// Parsing JSON Payload directly into a variable
		var item todo
		err := json.NewDecoder(res.Body).Decode(&item)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		fmt.Printf("%+v\n", item)
	}
}
