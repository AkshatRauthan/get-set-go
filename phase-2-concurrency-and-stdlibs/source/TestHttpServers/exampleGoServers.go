package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var waitFor int

func handler(w http.ResponseWriter, r *http.Request) {
	if waitFor > 0 {
		time.Sleep(time.Duration(waitFor) * time.Second)
	}
	fmt.Fprintf(w, "Hello, world! from %s\n", r.URL.Path[1:])
}

func ExampleGoServerForModule7(timeout int) {
	fmt.Printf("Running server with %dsec delay\n", timeout)

	if timeout > 0 {
		waitFor = timeout
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
