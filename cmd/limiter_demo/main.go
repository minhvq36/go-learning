package main

import (
	"fmt"
	"sync"

	"github.com/minhvq36/go-learning/internal/limiter"
)

func main() {

	const requestBuffer = 8
	requests := make(chan int, requestBuffer)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		limiter.ProcessRequest(requests)
	}()

	for req := range 5 {
		requests <- req
		fmt.Printf("[Main] Sent request %d\n", req)
	}
	close(requests)

	wg.Wait()
	fmt.Println("All requests processed!")
}
