package mock

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprintf(w, finalWord)
}
