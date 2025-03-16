package iteration

import "testing"

func TestIteration(t *testing.T) {

	t.Run("defaults to 5 repitions",
		func(t *testing.T) {
			repeated := Repeat("a", 5)
			expected := "aaaaa"

			if repeated != expected {
				t.Errorf("expected %q, but got %q", expected, repeated)
			}
		})

	t.Run("repeats 6 times when repeat number is specified",
		func(t *testing.T) {
			repeated := Repeat("a", 6)
			expected := "aaaaaa"

			if repeated != expected {
				t.Errorf("expected %q, but got %q", expected, repeated)
			}
		})

}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("x", 5)
	}
}
