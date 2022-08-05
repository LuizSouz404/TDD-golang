package main

import "testing"

func TestGreeting(t *testing.T) {
	verifyGreeting := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Greeting\nexpect: %v\ngot: %v", want, got)
		}
	}

	t.Run("This should returns 'Olá, Luiz' pass in args", func(t *testing.T) {
		got := Greeting("Luiz", "")
		want := "Olá, Luiz"

		verifyGreeting(t, got, want)
	})

	t.Run("This test should returns 'Olá, mundo' if args is empty", func(t *testing.T) {
		got := Greeting("", "")
		want := "Olá, mundo"

		verifyGreeting(t, got, want)
	})

	t.Run("This test should returns 'Hello, Luiz'", func(t *testing.T) {
		got := Greeting("Luiz", "english")
		want := "Hello, Luiz"

		verifyGreeting(t, got, want)
	})

	t.Run("This test should returns 'Bonjour, Luiz'", func(t *testing.T) {
		got := Greeting("Luiz", "french")
		want := "Bonjour, Luiz"

		verifyGreeting(t, got, want)
	})

}
