// di is acronym of Dependency Injection
package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	name := "Alireza"
	_ = Greet(&buffer, name)

	output := buffer.String()
	expectedOutput := EnglishHelloPrefix + name

	assertStringEqual(t, output, expectedOutput)
}

func assertStringEqual(t testing.TB, output string, expectedOutput string) {
	if output != expectedOutput {
		t.Errorf("want %q, got %q", expectedOutput, output)
	}
}
