package main

import (
	"log"
	"net/http"
	"time"

	"error-envelope-audit/server"
	"error-envelope-audit/store"
)

func main() {
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           server.NewHandler(store.New()),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}