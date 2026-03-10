package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, ch chan<- int) {
	defer close(ch)

	data := 42

	// mock time process
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		ch <- data
	case <-ctx.Done():
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan int)

	go fetchData(ctx, ch) // provide channel

	select {
	case val := <-ch:
		fmt.Printf("Received: %d", val)
	case <-ctx.Done():
		fmt.Println("Cancelled")
	}

}
