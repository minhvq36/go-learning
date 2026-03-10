package maps

import "fmt"

func RunExercise() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}

	scores["Carol"] = 75

	fmt.Println(scores)
}
