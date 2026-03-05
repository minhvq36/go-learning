package slice

import "fmt"

func RunExercise() {
	nums := []int{1, 2, 3}

	nums = append(nums, 4)

	fmt.Println(nums)
}
