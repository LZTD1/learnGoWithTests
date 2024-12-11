package sync

import "sync"

type Counter struct {
	value int
	m     sync.Mutex
}

func New() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
