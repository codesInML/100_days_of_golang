package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want, "test")
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound, "unknown")
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		want := "this is just a test"

		assertError(t, err, nil, "test")
		assertDefinition(t, dictionary, want, "test")
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExist, word)
		assertDefinition(t, dictionary, definition, word)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "new test"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil, word)
		assertDefinition(t, dictionary, newDefinition, word)
	})

	t.Run("update unexisting word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "new test"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist, word)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)

		assertDeleted(t, dictionary, word)
	})
}

func assertDeleted(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err == nil {
		t.Fatal("expected to get an error")
	}

	if got != "" {
		t.Errorf("expected %q but got %q, given %q", "", got, word)
	}
}

func assertStrings(t testing.TB, got, want, test string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q, given %q", want, got, test)
	}
}

func assertError(t testing.TB, got error, want error, test string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q, given %q", want, got, test)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, want, test string) {
	t.Helper()
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("expected to not get an error")
	}

	if got != want {
		t.Errorf("expected %q but got %q, given %q", want, got, test)
	}
}
