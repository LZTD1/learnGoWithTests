package memStorage

import "sync"

type MemStorage struct {
	store map[string]int
	m     sync.Mutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		store: make(map[string]int),
	}
}

func (m *MemStorage) GetPlayerScore(name string) int {
	return m.store[name]
}
func (m *MemStorage) RecordWin(name string) {
	m.m.Lock()
	m.store[name]++
	m.m.Unlock()
}
