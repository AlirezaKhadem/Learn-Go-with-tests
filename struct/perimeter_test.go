package struct_

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		width:  10.0,
		height: 10.0,
	}
	exceptedOutput := 40.0

	output := rectangle.Perimeter()
	if output != exceptedOutput {
		t.Errorf(
			"got %.2f want %.2f",
			output,
			exceptedOutput,
		)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			width:  6.0,
			height: 12.0,
		}
		exceptedOutput := 72.0

		output := rectangle.Area()
		if output != exceptedOutput {
			t.Errorf(
				"got %.2f want %.2f",
				output,
				exceptedOutput,
			)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{
			radius: 10,
		}
		exceptedOutput := 314.1592653589793
		output := circle.Area()

		if output != exceptedOutput {
			t.Errorf("got %g want %g", output, exceptedOutput)
		}
	})
}
