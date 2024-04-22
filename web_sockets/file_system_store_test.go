package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	const dbInitialData = `[
			{ "Name": "Cleo", "Wins": 10 },
			{ "Name": "Chris", "Wins": 33 }
		]`

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPLayerStore(database)

		AssertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, dbInitialData)
		defer cleanDatabase()

		store, err := NewFileSystemPLayerStore(database)

		AssertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScoreEquals(t, got, want)
	})

	t.Run("league from a render", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{ "Name": "Cleo", "Wins": 10 },
			{ "Name": "Chris", "Wins": 33 }
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPLayerStore(database)

		AssertNoError(t, err)
		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{ "Name": "Cleo", "Wins": 10 },
			{ "Name": "Chris", "Wins": 33 }
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPLayerStore(database)

		AssertNoError(t, err)
		got := store.GetPlayerScore("Chris")

		want := 33

		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, dbInitialData)
		defer cleanDatabase()

		store, err := NewFileSystemPLayerStore(database)

		AssertNoError(t, err)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPLayerStore(database)

		AssertNoError(t, err)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}