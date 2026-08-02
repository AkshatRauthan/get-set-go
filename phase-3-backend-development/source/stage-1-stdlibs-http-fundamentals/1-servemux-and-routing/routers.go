package main

import (
	"net/http"
	"phase-3/stage-1/1-servemux-and-routing/handlers"
)

// We are using a / req on top if matching happens on based of line declaration order
// That should catch all requests here.

// But in GO matching is performed on basis of how well defined the matching string is.....

// So the order of matching will be....
// -> GET /latest and GET /all
// -> GET /{messageId}
// -> /

func messageRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", handlers.GenericMessageHandler())
	mux.Handle("POST /", handlers.SendMessageHandler())

	mux.Handle("GET /all", handlers.GetAllMessagesHandler())
	mux.Handle("GET /{messageId}", handlers.ReceiveMessageHandler())
	mux.Handle("GET /latest", handlers.GetLatestMessageHandler())

	return mux
}
