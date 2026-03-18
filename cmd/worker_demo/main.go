package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/minhvq36/go-learning/internal/worker"
)

func main() {
	const numWorkers = 4
	const jobBuffer = numWorkers * 2
	var wg sync.WaitGroup

	jobs := make(chan int, jobBuffer)
	results := make(chan int, jobBuffer)

	// Start and close workers if jobs done
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker.StartWorker(workerID, jobs, results)
		}(w)
	}

	// WaitGroup monitoring goroutine to close results channel when all workers done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Add jobs
	go func() {
		for j := 0; j < 10; j++ {
			fmt.Printf("---> Main: Sending job %d to buffer\n", j)
			jobs <- j
			time.Sleep(time.Second)
		}

		close(jobs)
	}()

	// Turn out results from channel
	for res := range results {
		fmt.Printf("<--- Main: Received result %d\n", res)
	}

	fmt.Println("All jobs completed efficiently!")
}
