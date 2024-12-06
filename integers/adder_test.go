package integers

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			x, y int
		}
		want int
	}{
		{
			name:  "Test 1",
			input: struct{ x, y int }{x: 1, y: 1},
			want:  2,
		},
		{
			name:  "Test 2",
			input: struct{ x, y int }{x: -1, y: 1},
			want:  0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := Add(test.input.x, test.input.y)

			if sum != test.want {
				t.Errorf("Add(%d, %d) = %d, want %d", test.input.x, test.input.y, sum, test.want)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

	// Output: 6
}
