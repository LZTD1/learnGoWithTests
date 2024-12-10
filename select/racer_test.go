package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Default test", func(t *testing.T) {
		slowServer := createMockServer(20*time.Millisecond, t)
		fastServer := createMockServer(0*time.Millisecond, t)

		defer fastServer.Close()
		defer slowServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("Racer failed: %v", err)
		}
		if got != want {
			t.Errorf("Racer(slowURL, fastURL) = %s; want %s", got, want)
		}

	})
	t.Run("Test over 10s", func(t *testing.T) {
		server := createMockServer(50*time.Millisecond, t)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("Racer(server, server); want error")
		}
	})
}

func createMockServer(d time.Duration, t testing.TB) *httptest.Server {
	t.Helper()

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
