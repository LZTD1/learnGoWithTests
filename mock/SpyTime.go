package mock

import "time"

type SpyTime struct {
	duration time.Duration
}

func (s *SpyTime) Sleep(t time.Duration) {
	s.duration = t
}
