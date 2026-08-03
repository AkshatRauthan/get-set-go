package main

import (
	"fmt"
	"net/http"

	"phase-3/stage-1/2-request-parsing-and-response-envelope/handlers"
)

// -> /messages/  (subtree, stripped before reaching messageRouter)
// -> /           (subtree catch-all)

func Run() error {
	root := http.NewServeMux()

	root.Handle("/", handlers.GenericHandler())
	root.Handle("/messages/", http.StripPrefix("/messages", messageRouter()))

	server := &http.Server{
		Addr:    ":4040",
		Handler: root,
	}

	fmt.Println("Starting server on port 4040")
	return server.ListenAndServe()
}
