package main

import (
	"testing"
)

func TestSequence(t *testing.T) {

	t.Run("This test should return 'aaaa'", func(t *testing.T) {
		got := Sequence("", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("Sequence\nexpect: %v\ngot: %v", want, got)
		}
	})

	t.Run("This test should return 'bbbbb'", func(t *testing.T) {
		got := Sequence("b", 5)
		want := "bbbbb"

		if got != want {
			t.Errorf("Sequence\nexpect: %v\ngot: %v", want, got)
		}
	})

	t.Run("This test should return 'abcabc'", func(t *testing.T) {
		got := Sequence("abc", 2)
		want := "abcabc"

		if got != want {
			t.Errorf("Sequence\nexpect: %v\ngot: %v", want, got)
		}
	})
}
