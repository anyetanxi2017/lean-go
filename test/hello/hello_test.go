package main

import "testing"

func TestHello(t *testing.T) {
	// 分组测试

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("in Japan", func(t *testing.T) {
		got := Hello("yy", "Japan")
		want := "こんにちは,yy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("yy", "French")
		want := "Bonjour,yy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola,Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello,Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello,World"
		assertCorrectMessage(t, got, want)
	})
}
