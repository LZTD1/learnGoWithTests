package dependencyInjection

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "World!")

	got := buffer.String()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleGreet() {
	writer := os.Stdout
	Greet(writer, "World!")

	// Output: Hello World!
}
