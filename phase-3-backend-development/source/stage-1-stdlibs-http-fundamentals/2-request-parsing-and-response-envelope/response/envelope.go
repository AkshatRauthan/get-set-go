package response

import (
	"encoding/json"
	"net/http"
)

// Envelope is the single consistent JSON response shape used across every
// handler. omitempty means success responses won't show a stray "error":""
// and error responses won't show a stray "data":null.
type Envelope struct {
	Success bool   `json:"success"`
	Status  string `json:"status,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// Write is the low-level primitive every helper below funnels through.
func Write(w http.ResponseWriter, status int, envelope Envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope)
}

// Success writes a 2xx-style envelope with success=true and optional data.
func Success(w http.ResponseWriter, status int, statusMsg string, data any) {
	Write(w, status, Envelope{Success: true, Status: statusMsg, Data: data})
}

// Error writes a success=false envelope carrying only an error message.
func Error(w http.ResponseWriter, status int, errMsg string) {
	Write(w, status, Envelope{Success: false, Error: errMsg})
}
