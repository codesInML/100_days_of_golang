package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func (t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("rectangles", func (t *testing.T) {
		rectangle := Rectangle{4.0, 5.0}
		want := 18.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("circles", func (t *testing.T) {
		circle := Circle{5.6}
		want := 35.185837720205684
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {

	// checkArea := func (t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func (t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	want := 72.0
	// 	checkArea(t, rectangle, want)
	// })

	// t.Run("circles", func (t *testing.T) {
	// 	circle := Circle{10}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

	// Table Driven Test TDT
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` testname
		t.Run(tt.name, func (t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}