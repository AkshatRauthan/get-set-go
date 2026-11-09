package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"middleware-chaining/response"
)

// Recover turns a handler panic into a 500 instead of a dropped connection.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			rec := recover()
			if rec == nil {
				return
			}
			// Deliberate abort signal used by net/http; let it through.
			if rec == http.ErrAbortHandler {
				panic(rec)
			}
			log.Printf("panic: %v\n%s", rec, debug.Stack())
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}()
		next.ServeHTTP(w, r)
	})
}