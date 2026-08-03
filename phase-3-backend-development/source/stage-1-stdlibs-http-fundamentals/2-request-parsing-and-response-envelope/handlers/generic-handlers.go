package handlers

import (
	"fmt"
	"net/http"

	"phase-3/stage-1/2-request-parsing-and-response-envelope/response"
)

// genericNotFound collapses the two near-identical handlers you had before
// into one, parameterized by the path prefix used only for logging.
func genericNotFound(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Generic 404: %s %s%s\n", r.Method, prefix, r.URL.Path)
		response.Error(w, http.StatusNotFound,
			fmt.Sprintf("%s %s%s is not allowed", r.Method, prefix, r.URL.Path))
	}
}

func GenericHandler() http.HandlerFunc {
	return genericNotFound("")
}

func GenericMessageHandler() http.HandlerFunc {
	return genericNotFound("/messages")
}
