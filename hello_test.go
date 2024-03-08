package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Replit")
	expected := "Hello, Replit!"

	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}
