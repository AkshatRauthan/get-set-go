package server_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"error-envelope-audit/response"
	"error-envelope-audit/server"
	"error-envelope-audit/store"
)

// newTestHandler is the real handler served by main, plus a test-only route
// that panics so the Recover path is part of the audit.
func newTestHandler(t *testing.T) http.Handler {
	t.Helper()
	log.SetOutput(io.Discard) // Recover logs stack traces; keep test output clean.
	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	return server.NewHandler(store.New(), func(mux *http.ServeMux) {
		mux.HandleFunc("GET /panic", func(http.ResponseWriter, *http.Request) {
			panic("boom")
		})
	})
}

func do(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

// topLevelKeys returns the exact set of top-level JSON keys in body.
func topLevelKeys(t *testing.T, body []byte) map[string]bool {
	t.Helper()
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		t.Fatalf("body is not a JSON object: %v\nbody: %q", err, body)
	}
	keys := make(map[string]bool, len(raw))
	for k := range raw {
		keys[k] = true
	}
	return keys
}

// decodeStrict decodes into Envelope and rejects unknown fields and trailing data.
func decodeStrict(t *testing.T, body []byte) response.Envelope {
	t.Helper()
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	var env response.Envelope
	if err := dec.Decode(&env); err != nil {
		t.Fatalf("body does not match Envelope: %v\nbody: %q", err, body)
	}
	if dec.More() {
		t.Fatalf("trailing data after envelope: %q", body)
	}
	return env
}

func assertEqualKeys(t *testing.T, got map[string]bool, want ...string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("envelope keys = %v, want exactly %v", got, want)
	}
	for _, k := range want {
		if !got[k] {
			t.Fatalf("envelope keys = %v, missing %q", got, k)
		}
	}
}

// TestErrorEnvelopeAudit hits every registered route with a deliberately bad
// request, plus requests that never reach a route (mux 404/405), and asserts
// the body is always the error envelope.
func TestErrorEnvelopeAudit(t *testing.T) {
	h := newTestHandler(t)

	oversized := `{"text":"` + strings.Repeat("a", 1<<20+1) + `"}`

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		status int
	}{
		// POST /{$}
		{"send/malformed json", http.MethodPost, "/", `{"text":`, http.StatusBadRequest},
		{"send/empty body", http.MethodPost, "/", ``, http.StatusBadRequest},
		{"send/unknown field", http.MethodPost, "/", `{"text":"x","extra":1}`, http.StatusBadRequest},
		{"send/wrong type", http.MethodPost, "/", `{"text":123}`, http.StatusBadRequest},
		{"send/blank text", http.MethodPost, "/", `{"text":"   "}`, http.StatusBadRequest},
		{"send/body too large", http.MethodPost, "/", oversized, http.StatusRequestEntityTooLarge},
		{"send/wrong method", http.MethodGet, "/", ``, http.StatusMethodNotAllowed},

		// GET /receive/{messageId}
		{"receive/non-numeric id", http.MethodGet, "/receive/abc", ``, http.StatusBadRequest},
		{"receive/unknown id", http.MethodGet, "/receive/99", ``, http.StatusNotFound},
		{"receive/missing id", http.MethodGet, "/receive/", ``, http.StatusNotFound},
		{"receive/wrong method", http.MethodPost, "/receive/1", ``, http.StatusMethodNotAllowed},

		// GET /health: no exemption from the envelope
		{"health/wrong method post", http.MethodPost, "/health", ``, http.StatusMethodNotAllowed},
		{"health/wrong method delete", http.MethodDelete, "/health", ``, http.StatusMethodNotAllowed},

		// Recover path
		{"panic/recovered", http.MethodGet, "/panic", ``, http.StatusInternalServerError},

		// Never reaches a handler
		{"mux/unknown path", http.MethodGet, "/nope", ``, http.StatusNotFound},
		{"mux/unknown nested path", http.MethodPost, "/a/b/c", `{}`, http.StatusNotFound},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := do(h, tc.method, tc.path, tc.body)

			if rec.Code != tc.status {
				t.Fatalf("status = %d, want %d\nbody: %q", rec.Code, tc.status, rec.Body.String())
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Fatalf("Content-Type = %q, want application/json", ct)
			}
			if tc.status == http.StatusMethodNotAllowed && rec.Header().Get("Allow") == "" {
				t.Fatal("405 response is missing the Allow header")
			}

			body := rec.Body.Bytes()
			assertEqualKeys(t, topLevelKeys(t, body), "success", "error")

			env := decodeStrict(t, body)
			if env.Success {
				t.Error("success = true on an error response")
			}
			if env.Error == "" {
				t.Error("error is empty on an error response")
			}
			if env.Data != nil {
				t.Errorf("data = %v, want absent on an error response", env.Data)
			}
		})
	}
}

// TestSuccessEnvelopeAudit covers the other half of "no exceptions":
// successful responses, health check included, use the same envelope.
func TestSuccessEnvelopeAudit(t *testing.T) {
	h := newTestHandler(t)

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		status int
	}{
		{"health", http.MethodGet, "/health", ``, http.StatusOK},
		{"send", http.MethodPost, "/", `{"text":"AAPL long"}`, http.StatusCreated},
		{"receive", http.MethodGet, "/receive/1", ``, http.StatusOK}, // created by the previous row
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := do(h, tc.method, tc.path, tc.body)

			if rec.Code != tc.status {
				t.Fatalf("status = %d, want %d\nbody: %q", rec.Code, tc.status, rec.Body.String())
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Fatalf("Content-Type = %q, want application/json", ct)
			}

			body := rec.Body.Bytes()
			assertEqualKeys(t, topLevelKeys(t, body), "success", "data")

			env := decodeStrict(t, body)
			if !env.Success {
				t.Error("success = false on a success response")
			}
			if env.Error != "" {
				t.Errorf("error = %q, want absent on a success response", env.Error)
			}
			if env.Data == nil {
				t.Error("data is missing on a success response")
			}
		})
	}
}
