package daily

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

func SecondLargest(nums []int) (int, error) {
	/*
		- Find Second Largest in an Array
		- No sorting
		- No extra slice
		- Handle edge cases
	*/

	if len(nums) < 2 {
		err := fmt.Errorf("Input Array with fewer than 2 elements")
		return 0, err
	}

	largest, secondLargest := nums[0], math.MinInt
	foundSecondLargest := false

	for _, el := range nums[1:] {
		if el > largest {
			secondLargest = largest
			largest = el
			foundSecondLargest = true
		} else if el > secondLargest && el < largest {
			secondLargest = el
			foundSecondLargest = true
		}
	}

	if !foundSecondLargest {
		err := fmt.Errorf("No second largest element found")
		return 0, err
	}

	return secondLargest, nil
}

func GeneratePassword(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	/* Pre-allocate the byte slice for performance.
	 */
	password := make([]byte, n)
	charSetLen := big.NewInt(int64(len(letters)))

	for i := range n {
		/* Int returns a uniform random value in [0, max).
		 */
		idx, err := rand.Int(rand.Reader, charSetLen)
		if err != nil {
			return "", err
		}
		password[i] = letters[idx.Int64()]
	}

	return string(password), nil
}

type Limiter struct {
	mu       sync.Mutex
	requests map[string][]int64 // {key: []timeint} array of latest valid request time
	limit    int                // const set for how many request max for a window
	window   time.Duration      // const the length of window
}

func NewLimiter(limit int, window time.Duration) *Limiter {
	return &Limiter{
		requests: make(map[string][]int64), // init map for a new Limiter in memory
		limit:    limit,
		window:   window,
	}
}

func (limiter *Limiter) Allow(key string) bool {
	limiter.mu.Lock() // lock and unlock, allow only 1 request per time
	defer limiter.mu.Unlock()

	now := time.Now().Unix() // convert time to int64
	threshold := now - int64(limiter.window.Seconds())

	/*Clean expired stamps thoses are out of window*/
	var recentRequests []int64
	for _, stamp := range limiter.requests[key] {
		if stamp > threshold {
			recentRequests = append(recentRequests, stamp)
		}
	}

	if len(recentRequests) < limiter.limit {
		recentRequests = append(recentRequests, now) // add current request stamp that allowed
		limiter.requests[key] = recentRequests
		return true
	}
	limiter.requests[key] = recentRequests
	return false

}
