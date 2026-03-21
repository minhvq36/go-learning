package main

import (
	"fmt"

	"github.com/minhvq36/go-learning/internal/daily"
)

func main() {
	fmt.Println(daily.IsValidPrantheles("((("))
	if pass, err := daily.GeneratePassword(12); err == nil {
		fmt.Println("Generated Password:", pass)
	} else {
		fmt.Println("Error generating password:", err)
	}
}
