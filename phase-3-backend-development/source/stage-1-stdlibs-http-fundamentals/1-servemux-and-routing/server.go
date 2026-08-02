package main

import (
	"fmt"
	"net/http"
	"phase-3/stage-1/1-servemux-and-routing/handlers"
)

// We are using a / req on top if matching happens on based of line declaration order
// That should catch all requests here. [According to node's logic]

// But in GO matching is performed on basis of how well defined the matching string is.....

// So the order of matching will be....
// -> /messages/
// -> /

func Run() error {
	root := http.NewServeMux()

	root.Handle("/", handlers.GenericHandler())
	root.Handle("/messages/", http.StripPrefix("/messages", messageRouter()))
	// removing /messages from path so that the requests of the next mux can directly
	// catch their defined paths instead of adding /messages/ on every request string

	server := &http.Server{
		Addr:    ":4040",
		Handler: root,
	}

	fmt.Println("Starting server on port 4040")
	return server.ListenAndServe()
}
