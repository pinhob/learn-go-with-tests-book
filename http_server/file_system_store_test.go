package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a render", func(t *testing.T) {
		database := strings.NewReader(`[
			{ "Name": "Cleo", "Wins": 10 },
			{ "Name": "Chris", "Wins": 33 }
		]`)
		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{ "Name": "Cleo", "Wins": 10 },
			{ "Name": "Chris", "Wins": 33 }
		]`)
		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")

		want := 33

		assertScoreEqulas(t, got, want)
	})
}

func assertScoreEqulas(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
