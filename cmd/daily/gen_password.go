package main

import (
	"fmt"

	"github.com/minhvq36/go-learning/internal/daily"
)

func _() {
	password, err := daily.GeneratePassword(12)
	if err != nil {
		panic(err)
	}
	fmt.Println(password)
}
