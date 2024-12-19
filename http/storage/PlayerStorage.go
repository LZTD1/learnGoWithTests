package storage

type PlayerStorage interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}
