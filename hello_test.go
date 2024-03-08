package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("Replit", "")
		expected := "Hello, Replit"

		assertCorrectmessage(t, result, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World"

		assertCorrectmessage(t, result, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		result := Hello("Nachos", "Spanish")
		expected := "Hola, Nachos"

		assertCorrectmessage(t, result, expected)
	})

	t.Run("in French", func(t *testing.T) {
		result := Hello("Poirot", "French")
		expected := "Bonjour, Poirot"

		assertCorrectmessage(t, result, expected)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		result := Hello("Calabreso", "Portuguese")
		expected := "Oi, Calabreso"

		assertCorrectmessage(t, result, expected)
	})
}

func assertCorrectmessage(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}
