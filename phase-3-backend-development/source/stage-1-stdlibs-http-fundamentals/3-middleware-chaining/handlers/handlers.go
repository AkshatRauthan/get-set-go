package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"middleware-chaining/response"
	"middleware-chaining/store"
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

// PanicHandler exists only to exercise the Recover middleware.
func PanicHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	}
}