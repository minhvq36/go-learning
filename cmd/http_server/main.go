package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/minhvq36/go-learning/internal/api"
)

func main() {
	http.HandleFunc("/hello", api.HelloHandler)

	port := ":8080"
	fmt.Printf("Starting HTTP server on port %s...\n", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
