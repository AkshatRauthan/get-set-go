package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func CustomPutRequest() {
	url := "https://jsonplaceholder.typicode.com/users/1"

	payload := User{
		"John Doe",
		"a1@gmail.com",
		"1234567890",
		"a1.com",
		Company{
			"Firm A",
		},
		Address{
			"Montreal",
			"1233",
			Geo{
				"123",
				"123",
			},
		},
	}

	bodyBytes, err := json.Marshal(payload)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error marshalling payload: %v\n", err)
		return
	}
	client := http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		return
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		return
	}
	defer res.Body.Close()

	var user User

	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		fmt.Printf("%+v", user)
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing response: %v\n", err)
		return
	}

	fmt.Printf("Successfully posted user:\n %+v \n", user)
}
