package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/minhvq36/go-learning/internal/api"
)

func main() {
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("Starting server on :%d", port)
	log.Fatal(server.ListenAndServe())
}
