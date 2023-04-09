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

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "rectangle", shape: Rectangle{6.0, 12.0}, area: 72.0},
		{name: "circle", shape: Circle{10}, area: 314.1592653589793},
	}

	for _, testCase := range areaTests {
		checkArea(t, testCase.shape, testCase.area)
	}
}
