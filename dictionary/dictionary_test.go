package main

import "testing"

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{"test": "Test is just a test"}

	t.Run("This should search for a definition", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "Test is just a test"

		compareStrings(t, got, want)
	})

	t.Run("This should search but return an error", func(t *testing.T) {
		_, got := dictionary.Search("unknow")

		compareError(t, got, errorWordNotFound)
	})
}

func TestDictionaryAdd(t *testing.T) {
	t.Run("This should search for a definition", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "Test is just a test")

		got, _ := dictionary.Search("test")
		want := "Test is just a test"

		compareStrings(t, got, want)
	})
}

func compareError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Dictionary\ngot: %s\nexpect: %s", got, want)
	}

	if got == nil {
		t.Fatal("an error is expected to be obtained")
	}
}

func compareStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Dictionary\ngot: %s\nexpect: %s", got, want)
	}
}
