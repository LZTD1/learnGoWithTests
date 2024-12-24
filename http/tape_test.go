package http

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	f, c := createTempFile(t, "12345")
	defer c()

	tape := &tape{f}
	tape.Write([]byte("abc"))

	f.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(f)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
