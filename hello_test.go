package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Meera", "")
		want := "Hello, Meera"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to world when no name given", func(t *testing.T) {
		got := Hello("", "")

		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})
	t.Run("says hello in given language", func(t *testing.T) {
		got := Hello("Meera", "Spanish")

		want := "Hola, Meera"

		assertCorrectMessage(t, got, want)

	})
}
