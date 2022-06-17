package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

    t.Run("word is known", func (t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

        assertStrings(t, got, want)
    })
    t.Run("word is not known", func (t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := ErrNotFound

        assertErrors(t, err, want)
    })
}

func assertErrors(t testing.TB, got, want error) {
    t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
    t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}
