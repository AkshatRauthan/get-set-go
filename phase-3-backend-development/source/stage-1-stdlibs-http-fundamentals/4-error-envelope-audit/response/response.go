package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Envelope is the uniform JSON shape for every response.
type Envelope struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// Write encodes env as JSON with the given status code.
func Write(w http.ResponseWriter, status int, env Envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(env); err != nil {
		log.Printf("response: encode failed: %v", err)
	}
}

func Success(w http.ResponseWriter, status int, data any) {
	Write(w, status, Envelope{Success: true, Data: data})
}

func Error(w http.ResponseWriter, status int, msg string) {
	Write(w, status, Envelope{Success: false, Error: msg})
}