package struct_

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius int
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	return (rectangle.width + rectangle.height) * 2
}

func (rectangle Rectangle) Area() (area float64) {
	return rectangle.width * rectangle.height
}

func (circle Circle) Perimeter() (perimeter float64) {
	return 2.0 * math.Pi * float64(circle.radius)
}

func (circle Circle) Area() (area float64) {
	return math.Pi * math.Pow(float64(circle.radius), 2)
}
