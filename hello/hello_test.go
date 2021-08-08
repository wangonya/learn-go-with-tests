package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrtectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("sayting hello to people", func(t *testing.T) {
		got := Hello("You", "")
		want := "Hello, You!"
		assertCorrtectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrtectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrtectMessage(t, got, want)
	})
}
