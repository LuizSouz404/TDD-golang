package main

import (
	"fmt"
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

func BenchmarkSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sequence("abc", 1)
	}
}

func ExampleSequence() {
	sequence := Sequence("abc", 3)
	fmt.Println(sequence)
	// Output: abcabcabc
}
