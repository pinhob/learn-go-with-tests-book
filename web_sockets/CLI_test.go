package poker_test

import (
	"bytes"
	"fmt"
	poker "main/command-line"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {

		t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
			stdout := &bytes.Buffer{}
			in := strings.NewReader("Pies\n")
			game := &GameSpy{}

			cli := poker.NewCLI(in, stdout, game)
			cli.PlayPoker()

			if game.StartCalled {
				t.Errorf("game should not have started")
			}

			assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
		})

		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but god %d", game.StartedWith)
		}
	})

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")

		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("Expected finish called with 'Chris but got %q", game.FinishedWith)
		}
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")

		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("Expected finish called with 'Cleo but got %q", game.FinishedWith)
		}
	})
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
