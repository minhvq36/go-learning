package dto

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ApiError struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Details []FieldError `json:"details,omitempty"`
}

type ApiResponse[T any] struct {
	Success bool      `json:"success"`
	Data    T         `json:"data,omitempty"`
	Error   *ApiError `json:"error,omitempty"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type NoteResponse struct {
	NoteId  string `json:"note_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
