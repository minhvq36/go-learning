package maths

import (
	"errors"
	"fmt"
)

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("cannot divide by zero")
		return 0, fmt.Errorf("math error: %w", err)
	}
	return a / b, nil
}
