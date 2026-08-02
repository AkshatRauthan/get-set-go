package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GenericHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPath := r.URL.Path
		method := r.Method

		fmt.Printf("Generic Error 404: %s Req on %s\n", reqPath, method)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]any{"success": false, "status": "not found", "error": fmt.Sprintf("%s Req on %s is not allowed", method, reqPath)})
		return
	}
}

func GenericMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPath := r.URL.Path
		method := r.Method

		fmt.Printf("GenericMessage Error 404: %s Req on /messages/%s\n", reqPath, method)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]any{"success": false, "status": "not found", "error": fmt.Sprintf("%s Req on /messages/%s is not allowed", method, reqPath)})
		return
	}
}
