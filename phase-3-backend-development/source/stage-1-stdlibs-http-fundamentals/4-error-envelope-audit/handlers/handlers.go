package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"error-envelope-audit/response"
	"error-envelope-audit/store"
)

const maxBodyBytes = 1 << 20 // 1 MiB

type sendRequest struct {
	Text string `json:"text"`
}

func SendMessageHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var req sendRequest
		if err := dec.Decode(&req); err != nil {
			var tooBig *http.MaxBytesError
			if errors.As(err, &tooBig) {
				response.Error(w, http.StatusRequestEntityTooLarge, "request body too large")
				return
			}
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}
		if strings.TrimSpace(req.Text) == "" {
			response.Error(w, http.StatusBadRequest, "text is required")
			return
		}

		response.Success(w, http.StatusCreated, s.Add(req.Text))
	}
}

func ReceiveMessageHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("messageId"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "messageId must be an integer")
			return
		}
		m, ok := s.Get(id)
		if !ok {
			response.Error(w, http.StatusNotFound, "message not found")
			return
		}
		response.Success(w, http.StatusOK, m)
	}
}

// HealthHandler uses the envelope like every other route: no exceptions.
func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}
