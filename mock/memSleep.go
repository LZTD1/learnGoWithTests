package mock

type Sleeper interface {
	Sleep()
}

type MemSleeper struct {
	Calls int
}

func (s *MemSleeper) Sleep() {
	s.Calls++
}
