package api

import (
	"net/http"

	"github.com/minhvq36/go-learning/internal/api/dto"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		SendError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Only GET method is allowed")
		return
	}

	SendSuccess(w, http.StatusOK, dto.HelloResponse{Message: "Hello Go Backend with JSON"})
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		SendError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Only GET method is allowed")
		return
	}

	SendSuccess(w, http.StatusOK, dto.UserResponse{Id: 1, Name: "Minh"})
}
