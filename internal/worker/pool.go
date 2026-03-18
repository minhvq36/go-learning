package worker

import (
	"fmt"
	"time"
)

func StartWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d get job #%d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
		fmt.Printf("Worker %d done job #%d\n", id, j)
	}
}
