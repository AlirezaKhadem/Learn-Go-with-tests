package maps

import "testing"

func TestUpdate(t *testing.T) {
	t.Run("update", func(t *testing.T) {
		key := "key"
		value := "value"
		newValue := "new_value"
		dictionary := Dictionary{key: value}

		dictionary.Update(key, newValue)

		assertDefinition(t, dictionary, key, newValue)
	})
	t.Run("update not existing key", func(t *testing.T) {
		newKey := "new_key"
		value := "value"
		dictionary := Dictionary{}

		updateError := dictionary.Update(newKey, value)

		assertError(t, updateError, NotFoundError)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "key"
		value := "value"

		_ = dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing key", func(t *testing.T) {
		key := "key"
		value := "value"
		dictionary := Dictionary{key: value}

		addError := dictionary.Add(key, "another value")

		assertError(t, addError, KeyAlreadyExistsError)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "This is test value",
	}

	t.Run("know word", func(t *testing.T) {
		expectedOutput := "This is test value"
		output, _ := dictionary.Search("test")

		assertValue(t, output, expectedOutput)
	})
	t.Run("unknown word", func(t *testing.T) {
		value, searchError := dictionary.Search("umbrella")

		assertError(t, searchError, KeyDoesNotExistError)
		assertValue(t, value, "")
	})

}

func assertValue(t testing.TB, givenValue string, expectedValue string) {
	if givenValue != expectedValue {
		t.Errorf(
			"got %q, want %q",
			givenValue,
			expectedValue,
		)
	}
}

func assertError(t testing.TB, err error, expectedError error) {
	if err == nil {
		t.Error("error should not be nil. expected to get an error.")
	}
	if err != expectedError {
		t.Errorf(
			"want %q, got %q",
			expectedError.Error(),
			err.Error(),
		)
	}
}

func assertDefinition(
	t *testing.T,
	dictionary Dictionary,
	key string,
	value string,
) {
	t.Helper()

	output, searchErr := dictionary.Search(key)
	if searchErr != nil {
		t.Fatal("should find added word:", searchErr)
	}
	assertValue(t, output, value)
}
