package struct_

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, expectedOutput float64) {
		output := shape.Perimeter()

		if output != expectedOutput {
			t.Errorf(
				"got %.2f want %.2f",
				output,
				expectedOutput,
			)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			width:  10.0,
			height: 10.0,
		}
		expectedOutput := 40.0

		checkPerimeter(t, rectangle, expectedOutput)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{
			radius: 1,
		}
		expectedOutput := 2 * math.Pi

		checkPerimeter(t, circle, expectedOutput)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expectedOutput float64) {
		output := shape.Area()

		if output != expectedOutput {
			t.Errorf(
				"got %.2f want %.2f",
				output,
				expectedOutput,
			)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			width:  6.0,
			height: 12.0,
		}
		expectedOutput := 72.0

		checkArea(t, rectangle, expectedOutput)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{
			radius: 10,
		}
		expectedOutput := 314.1592653589793

		checkArea(t, circle, expectedOutput)
	})
}
