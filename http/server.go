package http

import (
	"fmt"
	"learnGoWithTests/http/storage"
	"net/http"
	"strings"
)

type PlayerServer struct {
	Store storage.PlayerStorage
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	player := strings.TrimPrefix(req.URL.Path, "/player/")
	switch req.Method {
	case http.MethodGet:
		p.processScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) processScore(w http.ResponseWriter, player string) {

	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
