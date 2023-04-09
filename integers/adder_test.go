package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)
	// Output: 5
}

func TestAdder(t *testing.T) {
	numberOne := 2
	numberTwo := 4
	output := Add(numberOne, numberTwo)

	assertCorrectMessage(t, numberOne, numberTwo, output)
}

func assertCorrectMessage(
	t testing.TB,
	numberOne int,
	numberTwo int,
	output int,
) {
	sum := numberOne + numberTwo
	if output != sum {
		t.Errorf(
			"sum of number %d and %d is %d, got %d",
			numberOne,
			numberTwo,
			sum,
			output,
		)
	}
}
