package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/minhvq36/go-learning/internal/api/dto"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path, e.g., /notes/{note_id}
	note_id := r.PathValue("note_id")

	title := "Note Title"
	message := fmt.Sprintf("Đang lấy dữ liệu cho Note ID: %s", note_id)
	SendSuccess(w, http.StatusOK, dto.NoteResponse{NoteId: note_id, Title: title, Content: message})
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateNoteRequest

	/* Block comment:
	   Use NewDecoder for streaming JSON directly from the request body.
	   This is more memory efficient than reading the entire body into a byte slice.
	*/
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "INVALID_JSON", "Cấu trúc JSON không hợp lệ")
		return
	}

	if req.Title == "" {
		SendError(w, http.StatusBadRequest, "REQUIRED_FIELD", "Tiêu đề không được để trống")
		return
	}

	log.Printf("Đã tạo Note: %s\n", req.Title)

	response := map[string]any{
		"note_id": "note_abc123",
		"title":   req.Title,
	}

	SendSuccess(w, http.StatusCreated, response)
}
