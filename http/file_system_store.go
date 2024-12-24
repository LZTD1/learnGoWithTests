package http

import (
	"encoding/json"
	"fmt"
	"io"
	"learnGoWithTests/http/entity"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   entity.League
}

func NewFSPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, io.SeekStart)
	err := initialPlayerDBFile(file)
	if err != nil {
		return nil, err
	}

	l, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{json.NewEncoder(&tape{file}), l}, nil
}

func initialPlayerDBFile(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file stat %s, %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func (f *FileSystemPlayerStore) GetLeague() entity.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
func (f *FileSystemPlayerStore) GetPlayerScore(n string) int {
	player := f.league.Find(n)
	if player != nil {
		return player.Wins
	}
	return 0
}
func (f *FileSystemPlayerStore) RecordWin(name string) {
	p := f.league.Find(name)

	if p != nil {
		p.Wins++
	} else {
		f.league = append(f.league, entity.Player{Name: name, Wins: 1})
	}

	f.database.Encode(f.league)
}
func NewLeague(rdr io.Reader) ([]entity.Player, error) {
	var l []entity.Player
	err := json.NewDecoder(rdr).Decode(&l)
	if err != nil {
		err = fmt.Errorf("Problems parsing league,  %v", err)
	}
	return l, err
}
