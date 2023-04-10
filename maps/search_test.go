package maps

import "testing"

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

		assertError(t, searchError, KeyNotExistError)
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
