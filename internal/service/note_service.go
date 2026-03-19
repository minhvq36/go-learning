package service

import (
	"context"
	"fmt"

	"github.com/minhvq36/go-learning/internal/repository"
)

func CreateNote(ctx context.Context, id int) error {
	fmt.Printf("[Service] Creating note %d...\n", id)
	err := repository.SaveNote(ctx, id)
	if err != nil {
		fmt.Printf("[Service] Failed to create note %d: %v\n", id, err)
		return fmt.Errorf("failed to create note: %w", err)
	} else {
		fmt.Printf("[Service] Successfully created note %d\n", id)
		return nil
	}
}
