package mock

import (
	"bytes"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Testing with MemSleeper", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		s := &MemSleeper{}

		Countdown(buffer, s)
		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if s.Calls != 3 {
			t.Errorf("got %d calls want 3", s.Calls)
		}
	})

	t.Run("Testing with fully mocking", func(t *testing.T) {
		m := &SpyCountdownOperations{}
		Countdown(m, m)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(m.Calls, want) {
			t.Errorf("got %q want %q", m.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	time := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{time, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.duration != time {
		t.Errorf("got %v want %v", spyTime.duration, time)
	}
}

func ExampleCountdown() {
	out := os.Stdout
	s := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(out, s)

	// Output: 3
	//2
	//1
	//Go!
}
