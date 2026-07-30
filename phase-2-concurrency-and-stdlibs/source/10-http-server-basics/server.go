package main

import (
	"http-server-basics/handlers"
	"http-server-basics/store"
	"net/http"
)

func Run() error {
	s := store.New()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /send", handlers.Send(s))
	mux.HandleFunc("GET /receive/{messageId}", handlers.Receive(s))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv.ListenAndServe()
}
