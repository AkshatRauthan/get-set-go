package middleware

import "net/http"

// Middleware wraps a handler with cross-cutting behavior.
type Middleware func(http.Handler) http.Handler

// Chain applies mws so that the first one listed is the outermost layer:
// Chain(h, A, B, C) == A(B(C(h))). Request flows A -> B -> C -> h,
// response unwinds C -> B -> A.
func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
