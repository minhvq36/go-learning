package main

import (
	"context"
	"fmt"
	"time"

	"github.com/minhvq36/go-learning/internal/service"
)

func main() {

	const timeout = 2
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	note_id := 321

	fmt.Printf("[Main] Creating note in %d seconds\n", timeout)
	err := service.CreateNote(ctx, note_id)

	if err != nil {
		fmt.Printf("[Main] Create note %d failed: %v\n", note_id, err)
	} else {
		fmt.Printf("[Main] Create note %d successfully\n", note_id)
	}
}
