package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		name := "Alireza"
		exceptedOutput := englishHelloPrefix + name
		output := Hello(name, english)

		assertCorrectMessage(t, output, exceptedOutput)
	})
	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		emptyString := ""
		exceptedOutput := englishHelloPrefix + "World"
		output := Hello(emptyString, english)

		assertCorrectMessage(t, output, exceptedOutput)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		name := "Elodie"
		exceptedOutput := spanishHelloPrefix + name
		output := Hello(name, spanish)

		assertCorrectMessage(t, output, exceptedOutput)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		name := "Mohammad"
		exceptedOutput := frenchHelloPrefix + name
		output := Hello(name, french)

		assertCorrectMessage(t, output, exceptedOutput)
	})
}

func assertCorrectMessage(
	t testing.TB,
	output string,
	exceptedOutput string,
) {
	t.Helper()
	if output != exceptedOutput {
		t.Errorf(
			"excepted %q got %q",
			exceptedOutput,
			output,
		)
	}
}
