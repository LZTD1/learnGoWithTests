package http

import (
	"fmt"
	"learnGoWithTests/http/storage/memStorage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWindAndRetrievingThem(t *testing.T) {
	store := memStorage.NewMemStorage()
	server := PlayerServer{Store: store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), postWinRequest(player))

	res := httptest.NewRecorder()
	server.ServeHTTP(res, getScoreRequest(player))
	assertsStatus(t, res.Code, http.StatusOK)
	assertsResponseBody(t, res.Body.String(), "3")

}
func getScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return req
}
func postWinRequest(n string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", n), nil)
	return req
}
