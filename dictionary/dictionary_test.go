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

		compareError(t, got, ErrorWordNotFound)
	})
}

func TestDictionaryAdd(t *testing.T) {
	t.Run("This should add a new definition", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "Test is just a test"
		dictionary.Add(word, definition)

		compareDefinition(t, dictionary, word, definition)
	})

	t.Run("This should failed on add request", func(t *testing.T) {
		word := "test"
		definition := "Test is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		compareError(t, err, ErrorWordExists)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestDictionaryUpdate(t *testing.T) {
	t.Run("This should update a definition", func(t *testing.T) {
		word := "test"
		definition := "Test is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "Update a test and a test is just a test"
		dictionary.Update(word, newDefinition)

		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("This should return a erro trying to update", func(t *testing.T) {
		word := "test"
		definition := "Test is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		compareError(t, err, ErrorWordNotExists)
	})
}
func TestDictionaryDelete(t *testing.T) {
	t.Run("This should delete a definition", func(t *testing.T) {
		word := "test"
		definition := "Test is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrorWordNotFound {
			t.Errorf("Dictionary\nexpect the word '%s' be deleted", word)
		}
	})
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatalf("should have found added word: %s", word)
	}

	if definition != got {
		t.Errorf("Dictionary\ngot: %s\nexpect: %s", got, definition)
	}
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
