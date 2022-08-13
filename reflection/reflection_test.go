package main

import (
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {
	t.Run("this test should have an value", func(t *testing.T) {
		want := "Chris"

		var result []string

		x := struct {
			Nome string
		}{want}

		through(x, func(input string) {
			result = append(result, input)
		})

		if result[0] != want {
			t.Errorf("Reflection\ngot: %s\nexpect: %s", result[0], want)
		}
	})
}

func through(x interface{}, fn func(input string)) {
	value := reflect.ValueOf(x)
	field := value.Field(0)
	fn(field.String())
}
