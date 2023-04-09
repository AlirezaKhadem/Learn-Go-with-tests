package main

import "testing"

func TestHello(t *testing.T) {
	name := "alireza"
	exceptedOutput := englishHelloPrefix + name

	output := Hello(name)
	if output != exceptedOutput {
		t.Errorf(
			"excepted %q got %q",
			exceptedOutput,
			output,
		)
	}
}