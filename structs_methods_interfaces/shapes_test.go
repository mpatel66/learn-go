package shapes

import "testing"

func TestRectangle(t *testing.T) {
	t.Run("returns perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}

	})

	t.Run("returns area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}
		got := Area(rectangle)
		want := 50.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}

	})
}
