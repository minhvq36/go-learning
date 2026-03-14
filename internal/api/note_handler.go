package api

import (
	"fmt"
	"net/http"

	"github.com/minhvq36/go-learning/internal/api/dto"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	message := fmt.Sprintf("Đang lấy dữ liệu cho Note ID: %s", id)
	SendSuccess(w, http.StatusOK, dto.NoteResponse{NoteId: id, Content: message})
}
