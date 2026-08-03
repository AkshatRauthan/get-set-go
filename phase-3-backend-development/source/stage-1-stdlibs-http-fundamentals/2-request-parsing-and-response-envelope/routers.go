package main

import (
	"net/http"

	"phase-3/stage-1/2-request-parsing-and-response-envelope/handlers"
)

// Matching precedence here:
// -> GET /all and GET /latest   (exact, most specific)
// -> POST /{$}                  (exact root only, thanks to {$})
// -> GET /{messageId}           (wildcard, single segment)
// -> /                          (subtree catch-all, least specific)

func messageRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", handlers.GenericMessageHandler())

	// {$} forces exact match on "/" only - without it, "POST /" was a subtree
	// match and would have swallowed POST requests to any nested path too.
	mux.Handle("POST /{$}", handlers.SendMessageHandler())

	mux.Handle("GET /all", handlers.GetAllMessagesHandler())
	mux.Handle("GET /latest", handlers.GetLatestMessageHandler())
	mux.Handle("GET /{messageId}", handlers.ReceiveMessageHandler())

	return mux
}
