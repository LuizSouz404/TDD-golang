package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	verifyFizzBuzzReturns := func(t *testing.T, result, waited string, value int) {
		t.Helper()
		if result != waited {
			t.Errorf("fizzbuzz(%d)\ngot: %v\nwant: %v", value, result, waited)
		}
	}

	t.Run("this test should return Fizz", func(t *testing.T) {
		value := 3
		got := fizzbuzz(value)
		want := "Fizz"

		verifyFizzBuzzReturns(t, got, want, value)
	})

	t.Run("this Test should return Buzz", func(t *testing.T) {
		value := 5
		got := fizzbuzz(value)
		want := "Buzz"

		verifyFizzBuzzReturns(t, got, want, value)
	})

	t.Run("this test should return FizzBuzz", func(t *testing.T) {
		value := 15
		got := fizzbuzz(value)
		want := "FizzBuzz"

		verifyFizzBuzzReturns(t, got, want, value)
	})

	t.Run("this test should return 4", func(t *testing.T) {
		value := 4
		got := fizzbuzz(value)
		want := "4"

		verifyFizzBuzzReturns(t, got, want, value)
	})

}
