package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Chris")

	got := buffer.String()
	want := "Olá, Chris"

	if got != want {
		t.Errorf("Greeting\ngot: %s\nwant: %s", got, want)
	}
}
