package struct_

import "testing"

func TestPerimeter(t *testing.T) {
	width, height := 10.0, 10.0
	exceptedOutput := 40.0

	output := Perimeter(width, height)
	if output != exceptedOutput {
		t.Errorf(
			"got %.2f want %.2f",
			output,
			exceptedOutput,
		)
	}
}

func TestArea(t *testing.T) {
	width, height := 12.0, 6.0
	exceptedOutput := 72.0

	output := Area(width, height)
	if output != exceptedOutput {
		t.Errorf(
			"got %.2f want %.2f",
			output,
			exceptedOutput,
		)
	}
}
