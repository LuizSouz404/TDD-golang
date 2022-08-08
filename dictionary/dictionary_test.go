package main

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("This should search for a definition", func(t *testing.T) {
		dictionary := map[string]string{"test": "Test is just a test"}

		got := Search(dictionary, "test")
		want := "Test is just a test"

		compareStrings(t, got, want)
	})
}

func compareStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Dictionary\ngot: %s\nexpect: %s", got, want)
	}
}
