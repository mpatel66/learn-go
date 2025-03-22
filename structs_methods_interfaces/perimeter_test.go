package structsmethodsinterfaces

import "testing"

func TestPeri(t *testing.T) {
	t.Run("returns perimeter of shape", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}

	})
}
