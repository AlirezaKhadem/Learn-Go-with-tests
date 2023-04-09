package struct_

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return (rectangle.width + rectangle.height) * 2
}

func Area(rectangle Rectangle) (area float64) {
	return rectangle.width * rectangle.height
}
