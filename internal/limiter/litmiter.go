package limiter

import (
	"fmt"
	"time"
)

func ProcessRequest(requests <-chan int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for req := range requests {
		<-ticker.C
		fmt.Printf("[Limiter] Processing request %d\n", req)
	}
}
