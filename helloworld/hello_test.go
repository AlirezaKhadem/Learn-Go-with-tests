package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		name := "Alireza"
		expectedOutput := englishHelloPrefix + name
		output := Hello(name, english)

		assertCorrectMessage(t, output, expectedOutput)
	})
	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		emptyString := ""
		expectedOutput := englishHelloPrefix + "World"
		output := Hello(emptyString, english)

		assertCorrectMessage(t, output, expectedOutput)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		name := "Elodie"
		expectedOutput := spanishHelloPrefix + name
		output := Hello(name, spanish)

		assertCorrectMessage(t, output, expectedOutput)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		name := "Mohammad"
		expectedOutput := frenchHelloPrefix + name
		output := Hello(name, french)

		assertCorrectMessage(t, output, expectedOutput)
	})
}

func assertCorrectMessage(
	t testing.TB,
	output string,
	expectedOutput string,
) {
	t.Helper()
	if output != expectedOutput {
		t.Errorf(
			"excepted %q got %q",
			expectedOutput,
			output,
		)
	}
}
