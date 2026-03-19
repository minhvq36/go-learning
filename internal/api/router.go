package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/minhvq36/go-learning/internal/api/dto"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		SendError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Only GET method is allowed")
		return
	}

	ctx := r.Context()

	fmt.Println("[HelloHandler] Start processing...")

	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("[HelloHandler] Done for %s\n", r.RemoteAddr)
		SendSuccess(w, http.StatusOK, dto.HelloResponse{Message: "Hello Go Backend with JSON"})

	case <-ctx.Done():
		fmt.Printf("[HelloHandler] Cancelled: %v\n", ctx.Err())
		// TODO: Send context to database layer to cancel ongoing DB queries if needed
		return
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		SendError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Only GET method is allowed")
		return
	}

	SendSuccess(w, http.StatusOK, dto.UserResponse{Id: 1, Name: "Minh"})
}
