package iterate

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	want := "aaaaaaaaaa"

	if repeated != want {
		t.Errorf("got %q, want %q", repeated, want)
	}
}

func ExampleRepeat() {
	s := "ba"
	r := Repeat("na", 2)
	fmt.Println(s + r)

	// Output: banana
}
