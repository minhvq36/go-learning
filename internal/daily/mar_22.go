package daily

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	index := len(*s) - 1
	element := (*s)[index]

	var zero T
	(*s)[index] = zero

	*s = (*s)[:index]
	return element, true
}

var dict = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func IsValidPrantheles(s string) bool {
	var stack Stack[rune]
	for _, char := range s {
		if next, open := dict[char]; open {
			stack.Push(next)
		} else {
			pop, ok := stack.Pop()
			if !ok || char != pop {
				return false
			}
		}
	}

	return len(stack) == 0
}
