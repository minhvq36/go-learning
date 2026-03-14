package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{Width: w, Height: h}
}

func (rect *Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func NewCircle(r float64) *Circle {
	return &Circle{Radius: r}
}

func (circ *Circle) Area() float64 { // use struct pointer for reduce memory copy if struct is large
	r := circ.Radius
	return math.Pi * r * r
}
