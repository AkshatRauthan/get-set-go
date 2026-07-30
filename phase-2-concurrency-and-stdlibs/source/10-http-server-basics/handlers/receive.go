package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"http-server-basics/store"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"
)

func Receive(s *store.Store) http.HandlerFunc {
	fmt.Println("Receive Request Received")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		w.Header().Set("Content-Type", "application/json")

		messageId := r.PathValue("messageId")
		messageId = strings.TrimSpace(messageId)

		if messageId == "" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "messageId is required"})
			return
		}

		// random between 3-7 seconds
		sleepDuration := rand.IntN(5) + 3

		select {

		case <-time.After(time.Duration(sleepDuration) * time.Second):
			msg, flag := s.Get(messageId)
			if flag {
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "message": msg})
				return
			}
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "requested message does not exist"})
			return

		case <-ctx.Done():
			w.WriteHeader(http.StatusRequestTimeout)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "request timed out"})
			return

		}
	}
}
