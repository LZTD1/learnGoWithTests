package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 3, 3, 1}, []int{5, 5})
	want := []int{10, 10}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Collection of any size", func(t *testing.T) {
		got := SumAllTails([]int{3, 3, 3, 1}, []int{5, 5})
		want := []int{7, 5}

		checkSums(t, got, want)
	})

	t.Run("Empty collection", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	x := []int{3, 3, 3, 1}
	y := []int{5, 5}

	got := SumAllTails(x, y)

	fmt.Println(got)
	// Output: [7 5]
}

func ExampleSum() {
	numbers := []int{3, 3, 3, 1}
	got := Sum(numbers)

	fmt.Println(got)
	// Output: 10
}

func ExampleSumAll() {
	x := []int{3, 3, 3, 1}
	y := []int{5, 5}

	got := SumAll(x, y)

	fmt.Println(got)
	// Output: [10 10]
}
