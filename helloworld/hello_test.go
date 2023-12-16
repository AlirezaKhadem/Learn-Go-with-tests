package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("alireza", "")
		want := "Hello, alireza!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world!' if when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("alireza", spanish)
		want := "Hola, alireza!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("gholi", french)
		want := "Bonjour, gholi!"

		assertCorrectMessage(t, got, want)
	})
}

func TestGetGreetingPrefix(t *testing.T) {
	t.Run("return correct corresponding prefix", func(t *testing.T) {
		got := getGreetingPrefix(french)
		want := frenchHelloPrefix

		assertCorrectMessage(t, got, want)
	})
	t.Run("return english prefix if supplied language are not valid", func(t *testing.T) {
		got := getGreetingPrefix("xxx")
		want := englishHelloPrefix

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
