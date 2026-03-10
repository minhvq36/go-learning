package main

import (
	"fmt"

	"github.com/minhvq36/go-learning/internal/maths"

	"github.com/minhvq36/go-learning/internal/slice"

	"github.com/minhvq36/go-learning/internal/maps"

	"github.com/minhvq36/go-learning/internal/models"
)

func main() {
	fmt.Println("Hello, Go")
	fmt.Println(maths.Sum(5, 3))
	slice.RunExercise()
	maps.RunExercise()
	user := models.NewUser("Alice", 25)
	models.Greet(user)
}
