package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
	Never forgot to add return after saving an error in Go...
	Else we can get two different responses like this...

	{
		"messages": [],
		"status": "messages list is empty",
		"success": true
	}
	{
		"message": {},
		"status": "received all messages successfully",
		"success": true
	}
*/

type Message struct {
	Message string `json:"message"`
}

var maxId = -1
var messages = make(map[int]string)

func SendMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send Message Request Received")
		w.Header().Set("Content-Type", "application/json")

		var msg Message
		maxId += 1
		msgIdx := maxId

		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid message format"})
			maxId -= 1
			return
		}

		messages[msgIdx] = msg.Message

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "status": "message send successfully", "message": messages[maxId]})
	}
}

func ReceiveMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Receive Message Request Received")
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("messageId")
		messageId, err := strconv.Atoi(id)

		if err == nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "error parsing message id"})
			return
		}

		if _, ok := messages[messageId]; !ok {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "message not found"})
			return
		}

		message := messages[messageId]
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "status": "message received successfully", "message": message})
	}
}

func GetAllMessagesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All Messages Request Received")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		if len(messages) == 0 {
			_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "status": "messages list is empty", "messages": []Message{}})
			return
		}

		_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "status": "received all messages successfully", "message": messages})
	}
}

func GetLatestMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get Latest Message Request Received")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		if maxId == -1 {
			_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "status": "no messages available for receiving", "message": Message{}})
			return
		}

		message := messages[maxId]
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "status": "latest message received successfully", "message": message})
	}
}
