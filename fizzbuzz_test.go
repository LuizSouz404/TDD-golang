package main

import "testing"

func TestFizzBuzz_3(t *testing.T) {
	got := fizzbuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(3)\ngot: %v\nwant: %v", got, want)
	}
}

func TestFizzBuzz_5(t *testing.T) {
	got := fizzbuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("fizzbuzz(buzz)\ngot: %v\nwant: %v", got, want)
	}
}

func TestFizzBuzz_15(t *testing.T) {
	got := fizzbuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(fizzbuzz)\ngot: %v\nwant: %v", got, want)
	}
}

func TestFizzBuzz_4(t *testing.T) {
	got := fizzbuzz(4)
	want := "4"

	if got != want {
		t.Errorf("fizzbuzz(4)\ngot: %v\nwant: %v", got, want)
	}
}
