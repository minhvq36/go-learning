package api

import "net/http"

/* Block comment:
   RegisterRoutes accepts a pointer to ServeMux to avoid memory copying
   and ensure routes are registered to the original instance.
*/
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /hello", HelloHandler)
	mux.HandleFunc("GET /user", UserHandler)
	mux.HandleFunc("GET /notes/{note_id}", GetNoteHandler)
	mux.HandleFunc("POST /notes", CreateNoteHandler)
}

// TODO: Split routes into separates for better organization (e.g., user_routes.go, note_routes.go)
