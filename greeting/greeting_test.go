package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "Olá, Luiz"

	if got != want {
		t.Errorf("Greeting\nexpect: %v\ngot: %v", want, got)
	}
}
