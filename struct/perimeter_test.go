package struct_

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		width:  10.0,
		height: 10.0,
	}
	exceptedOutput := 40.0

	output := Perimeter(rectangle)
	if output != exceptedOutput {
		t.Errorf(
			"got %.2f want %.2f",
			output,
			exceptedOutput,
		)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{
		width:  6.0,
		height: 12.0,
	}
	exceptedOutput := 72.0

	output := Area(rectangle)
	if output != exceptedOutput {
		t.Errorf(
			"got %.2f want %.2f",
			output,
			exceptedOutput,
		)
	}
}
