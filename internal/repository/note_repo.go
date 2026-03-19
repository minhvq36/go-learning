package repository

import (
	"context"
	"fmt"
	"time"
)

func SaveNote(ctx context.Context, id int) error {
	mockTimer := time.NewTimer(3 * time.Second)
	defer mockTimer.Stop()
	select {
	case <-mockTimer.C:
		fmt.Printf("[Repo] Successfully saved note %d\n", id)
		return nil
	case <-ctx.Done():
		fmt.Printf("[Repo] Cancelled saving note %d: %v\n", id, ctx.Err())
		return ctx.Err()
	}
}
