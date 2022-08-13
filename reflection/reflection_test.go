package main

import "testing"

func TestReflection(t *testing.T) {
	want := "Chris"

	var result []string

	x := struct {
		Nome string
	}{want}

	through(x, func(input string) {
		result = append(result, input)
	})

	if len(result) != 1 {
		t.Errorf("Reflection\ngot: %d\nexpect: %d", len(result), 1)
	}
}

func through(x interface{}, fn func(input string)) {
	fn("I still not believe then Brazil lost for 7x1")
}
