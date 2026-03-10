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
	user.Greet()

	fmt.Printf("Age before: %d\n", user.Age)
	user.IncreaseAge()
	fmt.Printf("Age after: %d\n", user.Age)

	result, err := maths.Divide(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

}
