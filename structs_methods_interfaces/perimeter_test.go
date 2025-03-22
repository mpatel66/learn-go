package structsmethodsinterfaces

import "testing"

func TestRectangle(t *testing.T) {
	t.Run("returns perimeter of rectangle", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}

	})

	t.Run("returns area of rectangle", func(t *testing.T) {
		got := Area(5.0, 10.0)
		want := 50.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}

	})
}
