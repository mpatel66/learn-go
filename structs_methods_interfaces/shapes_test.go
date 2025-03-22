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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{5, 10}, 50.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %.g, want %.g", got, tt.want)
		}
	}

}
