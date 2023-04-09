package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat a character 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expectedOutput := "aaaaa"

		if repeated != expectedOutput {
			t.Errorf("expected %q but got %q", expectedOutput, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for index := 0; index <= b.N; index++ {
		Repeat("a")
	}
}
