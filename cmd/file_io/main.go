package main

import (
	"fmt"
	"log"

	"github.com/minhvq36/go-learning/internal/storage"
)

func main() {
	fileName := "note.txt"
	content := "Content save test"

	err := storage.SaveNote(fileName, content)

	if err != nil {
		log.Fatalf("Error saving note: %v\n", err)
	} else {
		fmt.Println("Note saved successfully!")
	}

	content, err = storage.LoadNote(fileName)

	if err != nil {
		log.Fatalf("Error loading note: %v\n", err)
	} else {
		fmt.Printf("Content: %s\n", content)
	}
}
