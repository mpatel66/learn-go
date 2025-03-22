package shapes

import "testing"

type Shape interface {
	Area() float64
}

func TestRectangle(t *testing.T) {
	t.Run("returns perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}

	})

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.g, want %.g", got, want)
		}

	}
	t.Run("returns area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}
		want := 50.0

		checkArea(t, rectangle, want)

	})

	t.Run("returns are of circle", func(t *testing.T) {
		circle := Circle{10}

		want := 314.1592653589793
		checkArea(t, circle, want)

	})
}
