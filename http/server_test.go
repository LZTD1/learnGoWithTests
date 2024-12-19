package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStorage struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStorage) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStorage) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGetPlayers(t *testing.T) {
	storage := StubPlayerStorage{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	s := &PlayerServer{Store: &storage}

	t.Run("returns Pepper`s score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)

		assertsStatus(t, res.Code, http.StatusOK)
		assertsResponseBody(t, res.Body.String(), "20")
	})

	t.Run("returns Floyd`s score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)

		assertsResponseBody(t, res.Body.String(), "10")
		assertsStatus(t, res.Code, http.StatusOK)
	})
	t.Run("Returns 404 on missing player", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)

		assertsStatus(t, res.Code, http.StatusNotFound)
	})
}
func TestStoreWins(t *testing.T) {
	storage := StubPlayerStorage{
		scores: map[string]int{},
	}
	s := &PlayerServer{Store: &storage}

	t.Run("it returns accepted in POST", func(t *testing.T) {
		player := "Pepper"
		req, _ := newPostWinRequest(player)
		res := httptest.NewRecorder()

		s.ServeHTTP(res, req)

		assertsStatus(t, res.Code, http.StatusAccepted)
		if len(storage.winCalls) != 1 {
			t.Errorf("expecting 1 call to store win, got %d", len(storage.winCalls))
		}
		if storage.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", storage.winCalls[0], player)
		}
	})
}

func newPostWinRequest(n string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", n), nil)
}

func assertsStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return req
}

func assertsResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
