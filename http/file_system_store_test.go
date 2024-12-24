package http

import (
	"learnGoWithTests/http/entity"
	"log"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("League from a reader", func(t *testing.T) {
		db, clean := createTempFile(t, `[
			{"Name": "Chris", "Wins": 10},
			{"Name": "Cleo", "Wins": 33}]`)
		defer clean()

		store, err := NewFSPlayerStore(db)
		if err != nil {
			log.Fatalf("problem creating file system player store, %v ", err)
		}
		got := store.GetLeague()

		want := []entity.Player{
			{"Cleo", 33},
			{"Chris", 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, clean := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer clean()

		store, err := NewFSPlayerStore(db)
		if err != nil {
			log.Fatalf("problem creating file system player store, %v ", err)
		}
		got := store.GetPlayerScore("Chris")
		assertScoreEquals(t, got, 33)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		db, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFSPlayerStore(db)
		if err != nil {
			log.Fatalf("problem creating file system player store, %v ", err)
		}
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		db, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFSPlayerStore(db)
		if err != nil {
			log.Fatalf("problem creating file system player store, %v ", err)
		}

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFSPlayerStore(database)

		assertNoError(t, err)
	})

}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got %v, wanted no error", err)
	}
}

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmp, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatal(err)
	}

	tmp.Write([]byte(initialData))

	removeFile := func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}

	return tmp, removeFile
}
