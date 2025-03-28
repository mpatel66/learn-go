package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get error.")
		}

		assertError(t, got, ErrNotFound)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {

		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {

		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("add a new word", func(t *testing.T) {
		definition := "this is just a test"
		word := "test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "apple"
		newDefinition := "a round, red fruit"
		definition := "a round, green fruit"

		dictionary := Dictionary{"apple": definition}

		err := dictionary.Add(word, newDefinition)

		assertError(t, err, ErrWordExists)

		assertDefinition(t, dictionary, word, definition)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("updates existing word", func(t *testing.T) {
		word := "apple"
		newDefinition := "a round, red fruit"
		oldDefinition := "a round, green fruit"

		dictionary := Dictionary{"apple": oldDefinition}
		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("does not add new word", func(t *testing.T) {
		word := "apple"
		definition := "a round, red fruit"

		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T) {
	t.Run("updates existing word", func(t *testing.T) {
		word := "apple"
		oldDefinition := "a round, green fruit"

		dictionary := Dictionary{"apple": oldDefinition}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)

	})

}
