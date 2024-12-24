package storage

import "learnGoWithTests/http/entity"

type PlayerStorage interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() entity.League
}
