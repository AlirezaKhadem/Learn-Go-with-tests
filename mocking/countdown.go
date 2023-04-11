package mocking

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer, sleeper Sleeper) error {
	for index := 3; index >= 1; index-- {
		_, fprintlnErro := fmt.Fprintln(writer, index)
		sleeper.Sleep()

		if fprintlnErro != nil {
			return fprintlnErro
		}
	}
	_, fprintError := fmt.Fprint(writer, "Go!")
	if fprintError != nil {
		return fprintError
	}
	return nil
}
