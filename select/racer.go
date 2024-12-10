package _select

import (
	"errors"
	"net/http"
	"time"
)

func ConfigurableRacer(a, b string, t time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(t):
		return "", errors.New("timeout")
	}
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ping(u string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(u)
		close(ch)
	}()
	return ch
}
