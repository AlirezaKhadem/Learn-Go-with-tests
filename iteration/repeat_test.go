package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat a character 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		exceptedOutput := "aaaaa"

		if repeated != exceptedOutput {
			t.Errorf("expected %q but got %q", exceptedOutput, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for index := 0; index <= b.N; index++ {
		Repeat("a")
	}
}
