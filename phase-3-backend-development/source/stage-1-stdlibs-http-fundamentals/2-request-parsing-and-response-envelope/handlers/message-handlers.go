package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"phase-3/stage-1/2-request-parsing-and-response-envelope/response"
)

type Message struct {
	Message string `json:"message"`
}

// mu guards maxId and messages, since both are shared mutable state hit
// concurrently by every request.
var (
	mu       sync.RWMutex
	maxId    = -1
	messages = make(map[int]string)
)

const maxBodyBytes = 1 << 20 // 1MB cap on incoming request bodies

func SendMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send Message Request Received")

		// Cross-cutting, applies to every route the same way - fine to do
		// inline here rather than wait for the middleware chain topic.
		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

		var msg Message
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // reject unexpected fields instead of silently ignoring them

		if err := decoder.Decode(&msg); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid message format")
			return
		}

		// Decode succeeding doesn't mean the field was meaningfully provided -
		// {} decodes fine into Message{Message: ""}. Reject that explicitly.
		if strings.TrimSpace(msg.Message) == "" {
			response.Error(w, http.StatusBadRequest, "message field cannot be empty")
			return
		}

		mu.Lock()
		maxId++ // only incremented after validation succeeds - no more burned IDs on bad input
		msgIdx := maxId
		messages[msgIdx] = msg.Message
		mu.Unlock()

		response.Success(w, http.StatusOK, "message sent successfully", map[string]any{
			"messageId": msgIdx,
			"message":   msg.Message,
		})
	}
}

func ReceiveMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Receive Message Request Received")

		idStr := r.PathValue("messageId")
		id, err := strconv.Atoi(idStr)
		if err != nil || id < 0 {
			response.Error(w, http.StatusBadRequest, "invalid messageId")
			return
		}

		mu.RLock()
		message, ok := messages[id]
		mu.RUnlock()

		if !ok {
			response.Error(w, http.StatusNotFound, "message not found")
			return
		}

		response.Success(w, http.StatusOK, "message received successfully", map[string]any{
			"messageId": id,
			"message":   message,
		})
	}
}

func GetAllMessagesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All Messages Request Received")

		// GET /all?limit=N - first real query param usage
		limit := -1
		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			parsedLimit, err := strconv.Atoi(limitStr)
			if err != nil || parsedLimit < 0 {
				response.Error(w, http.StatusBadRequest, "invalid limit query param")
				return
			}
			limit = parsedLimit
		}

		mu.RLock()
		defer mu.RUnlock()

		if len(messages) == 0 {
			response.Success(w, http.StatusOK, "messages list is empty", map[int]string{})
			return
		}

		if limit == -1 || limit >= len(messages) {
			response.Success(w, http.StatusOK, "received all messages successfully", messages)
			return
		}

		limited := make(map[int]string, limit)
		count := 0
		for id, m := range messages {
			if count >= limit {
				break
			}
			limited[id] = m
			count++
		}
		response.Success(w, http.StatusOK, "received limited messages successfully", limited)
	}
}

func GetLatestMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get Latest Message Request Received")

		mu.RLock()
		defer mu.RUnlock()

		if maxId == -1 {
			response.Success(w, http.StatusOK, "no messages available for receiving", nil)
			return
		}

		response.Success(w, http.StatusOK, "latest message received successfully", map[string]any{
			"messageId": maxId,
			"message":   messages[maxId],
		})
	}
}
