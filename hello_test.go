package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		name := "alireza"
		exceptedOutput := englishHelloPrefix + name
		output := Hello(name, englishLanguage)

		assertCorrectMessage(t, output, exceptedOutput)
	})
	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		emptyString := ""
		exceptedOutput := englishHelloPrefix + "World"
		output := Hello(emptyString, englishLanguage)

		assertCorrectMessage(t, output, exceptedOutput)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		name := "Elodie"
		exceptedOutput := spanishHelloPrefix + name
		output := Hello(name, spanishLanguage)

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
