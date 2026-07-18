package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	CompanyName string `json:"name"`
}

type User struct {
	UserName string  `json:"name"`
	Email    string  `json:"email"`
	PhoneNo  string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
	Address  Address `json:"address"`
}

// With the below syntax we can only send GET, HEAD or POST request.
// To use another kind of requests like PATCH, DELETE or PUT we use a user-defined request

func ApiReqWithNestedJsonResponse() {
	var url = "https://jsonplaceholder.typicode.com/users/1"
	client := http.Client{Timeout: 1 * time.Second}
	res, err := client.Get(url)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(-1)
	}
	defer res.Body.Close()

	var nestedJson User

	err = json.NewDecoder(res.Body).Decode(&nestedJson)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Printf("%+v\n", nestedJson)
}
