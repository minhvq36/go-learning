package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/minhvq36/go-learning/internal/daily"
)

const (
	MaxRequests = 5
	WindowTime  = 10 * time.Second
)

func main() {
	nums := []int{5, 5}
	secondLargest, err := daily.SecondLargest(nums)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Second Largest:", secondLargest)

	limiter := daily.NewLimiter(MaxRequests, WindowTime)
	userID := "random_user_007"

	fmt.Printf("--- 🛡️  Rate Limiter Test: %d reqs / %v ---\n", MaxRequests, WindowTime)
	fmt.Println("Kịch bản: User gửi 15 requests với khoảng nghỉ ngẫu nhiên (0.5s - 3s)")
	fmt.Println("------------------------------------------------------------")

	/* Simulate 15 requests with random intervals
	 */
	for i := 1; i <= 15; i++ {
		now := time.Now()
		allowed := limiter.Allow(userID)

		status := "✅ ALLOWED"
		if !allowed {
			status = "❌ DENIED "
		}

		fmt.Printf("[%s] Req %2d: %s\n", now.Format("15:04:05"), i, status)

		/* Random sleep between 500ms and 3000ms
		   Giúp mô phỏng hành vi User thực tế hơn là dùng 100ms cố định.
		*/
		sleepTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond
		time.Sleep(sleepTime)
	}
}
