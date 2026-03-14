package main

import (
	"fmt"

	"github.com/minhvq36/go-learning/internal/shapes"
)

func printArea(s shapes.Shape) {
	fmt.Printf("%.2f\n", s.Area())
}

func main() {
	rect := shapes.NewRectangle(5, 6)
	circ := shapes.NewCircle(3)

	printArea(rect)
	printArea(circ)
}
