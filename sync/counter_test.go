package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment 3 times", func(t *testing.T) {
		c := New()
		c.Inc()
		c.Inc()
		c.Inc()

		assetCounter(t, c, 3)
	})
	t.Run("Increment gorutines", func(t *testing.T) {
		want := 1000
		c := New()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assetCounter(t, c, want)
	})
}

func assetCounter(t *testing.T, c *Counter, n int) {
	if c.Value() != n {
		t.Errorf("Wanted %d, but got %d", 3, c.Value())
	}
}
