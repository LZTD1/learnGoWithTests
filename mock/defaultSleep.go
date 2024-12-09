package mock

import "time"

type DefaultSleep struct {
}

func (d *DefaultSleep) Sleep() {
	time.Sleep(1 * time.Second)
}
