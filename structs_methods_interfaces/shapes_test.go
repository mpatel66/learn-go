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
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{5, 10}, hasArea: 50.0},
		{name: "circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g, but has area %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}
