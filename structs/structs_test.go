package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{
		Width:  10.,
		Height: 10.,
	}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {
	areaStruct := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12., 6.}, 72.},
		{"Circle", Circle{10.}, 314.1592653589793},
		{"Circle", Triangle{12, 6}, 36},
	}

	for _, s := range areaStruct {
		t.Run(s.name, func(t *testing.T) {
			got := s.shape.Area()
			if got != s.want {
				t.Errorf("%#v got %g want %g", s.shape, got, s.want)
			}
		})
	}
}
