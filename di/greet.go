package di

import (
	"fmt"
	"io"
)

const (
	EnglishHelloPrefix = "Hello, "
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprint(writer, EnglishHelloPrefix+name)
	if err != nil {
		return err
	}
	return nil
}
