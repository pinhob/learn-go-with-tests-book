package smi

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	/*
		  or this way, trying inidivdually with a helper function:
			checkArea := func(t testing.TB, shape Shape, want float64) {
				got := shape.Area()
				if got != want {
		      t.Errorf("got %g want %g", got, want)
				}
			}

			t.Run("rectangles", func(t *testing.T) {
				rectangle := Rectangle{10.0, 10.0}
				checkArea(t, rectangle, 100.0)
			})

			t.Run("circles", func(t *testing.T) {
				circle := Circle{10}
				checkArea(t, circle, 314.1592653589793)
			})
	*/
}
