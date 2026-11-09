// Package server builds the full HTTP handler (routes + middleware) so that
// main and the tests exercise exactly the same thing.
package server

import (
	"net/http"

	"error-envelope-audit/store"
	"error-envelope-audit/handlers"
	"error-envelope-audit/response"
	"error-envelope-audit/middlewares"
)

// NewHandler returns the routed, middleware-wrapped handler. extra lets
// callers (tests) register additional routes on the mux before it is sealed.
func NewHandler(s *store.Store, extra ...func(*http.ServeMux)) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("POST /{$}", handlers.SendMessageHandler(s))
	mux.Handle("GET /receive/{messageId}", handlers.ReceiveMessageHandler(s))
	mux.Handle("GET /health", handlers.HealthHandler())
	for _, register := range extra {
		register(mux)
	}

	// Logging outermost so it observes the 500 that Recover writes.
	return middleware.Chain(envelopeMuxErrors(mux), middleware.Logging, middleware.Recover)
}

// envelopeMuxErrors replaces ServeMux's plain-text 404/405 replies with the
// envelope. Those replies are produced by the mux itself, before any of our
// handlers run, so they can't be fixed inside a handler.
func envelopeMuxErrors(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, pattern := mux.Handler(r)
		if pattern == "" {
			probe := &probeWriter{header: http.Header{}}
			h.ServeHTTP(probe, r)
			if probe.status == http.StatusNotFound || probe.status == http.StatusMethodNotAllowed {
				if allow := probe.header.Get("Allow"); allow != "" {
					w.Header().Set("Allow", allow)
				}
				response.Error(w, probe.status, http.StatusText(probe.status))
				return
			}
			// Anything else (e.g. path-cleaning redirect) is left to the mux.
		}
		mux.ServeHTTP(w, r)
	})
}

// probeWriter records what the mux's built-in handler would have sent and
// discards the body.
type probeWriter struct {
	header http.Header
	status int
}

func (p *probeWriter) Header() http.Header { return p.header }

func (p *probeWriter) WriteHeader(code int) {
	if p.status == 0 {
		p.status = code
	}
}

func (p *probeWriter) Write(b []byte) (int, error) {
	if p.status == 0 {
		p.status = http.StatusOK
	}
	return len(b), nil
}
