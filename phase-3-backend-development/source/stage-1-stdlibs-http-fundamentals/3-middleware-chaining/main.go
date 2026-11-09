package main

import (
	"log"
	"net/http"
	"time"

	"middleware-chaining/store"
	"middleware-chaining/handlers"
	"middleware-chaining/middlewares"
)

func main() {
	s := store.New()

	mux := http.NewServeMux()
	mux.Handle("POST /{$}", handlers.SendMessageHandler(s))
	mux.Handle("GET /receive/{messageId}", handlers.ReceiveMessageHandler(s))
	mux.Handle("GET /panic", handlers.PanicHandler())

	// Logging outermost so it observes the 500 that Recover writes.
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           middleware.Chain(mux, middleware.Logging, middleware.Recover),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}