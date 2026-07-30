package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"http-server-basics/store"
	"math/rand/v2"
	"net/http"
	"time"
)

func Send(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send Request Received")
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		w.Header().Set("Content-Type", "application/json")

		var msg store.Message
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid message format"})
		}

		// random between 3-7 seconds
		sleepDuration := rand.IntN(5) + 3

		select {

		case <-time.After(time.Duration(sleepDuration) * time.Second):
			s.Put(&msg)
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "message": "message sent successfully"})
			return

		case <-ctx.Done():
			w.WriteHeader(http.StatusRequestTimeout)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "request timed out"})
			return
		}
	}
}
