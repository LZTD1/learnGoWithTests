package http

import (
	"encoding/json"
	"fmt"
	"io"
	"learnGoWithTests/http/entity"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStorage struct {
	scores   map[string]int
	winCalls []string
	league   []entity.Player
}

func (s *StubPlayerStorage) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStorage) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func (s *StubPlayerStorage) GetLeague() entity.League {
	return s.league
}

func TestGetPlayers(t *testing.T) {
	storage := StubPlayerStorage{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	s := NewPlayerServer(&storage)

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
	s := NewPlayerServer(&storage)

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
func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []entity.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStorage{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertsStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
	})
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("expecting application/json, got %s", response.Result().Header.Get("Content-Type"))
	}
}

func assertLeague(t *testing.T, got []entity.Player, league []entity.Player) {
	t.Helper()

	if reflect.DeepEqual(got, league) == false {
		t.Fatalf("Not Equals! %+v, want %+v", got, league)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (l []entity.Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&l)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return l
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newPostWinRequest(n string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", n), nil)
}

func assertsStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertsResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
