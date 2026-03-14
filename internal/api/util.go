package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/minhvq36/go-learning/internal/api/dto"
)

func sendJSON[T any](w http.ResponseWriter, status int, body dto.ApiResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("ERROR: Failed to encode JSON response: %v", err)
	}

}

func SendSuccess[T any](w http.ResponseWriter, status int, data T) {
	sendJSON(w, status, dto.ApiResponse[T]{
		Success: true,
		Data:    data,
	})
}

func SendError(w http.ResponseWriter, status int, code, message string) {
	sendJSON(w, status, dto.ApiResponse[any]{
		Success: false,
		Error:   &dto.ApiError{Code: code, Message: message},
	})
}
