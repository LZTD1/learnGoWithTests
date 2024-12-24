package http

import (
	"fmt"
	"learnGoWithTests/http/entity"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWindAndRetrievingThem(t *testing.T) {
	db, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := NewFSPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))

	t.Run("Get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, getScoreRequest(player))
		assertsStatus(t, res.Code, http.StatusOK)
		assertsResponseBody(t, res.Body.String(), "3")
	})
	t.Run("Get League", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())
		assertsStatus(t, res.Code, http.StatusOK)

		got := getLeagueFromResponse(t, res.Body)
		want := []entity.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

}
func getScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
func postWinRequest(n string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", n), nil)
	return req
}
