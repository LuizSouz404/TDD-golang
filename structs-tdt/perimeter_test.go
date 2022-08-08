package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("this should returns 40.0 perimeter", func(t *testing.T) {
		retangule := Retangule{10.0, 10.0}
		got := retangule.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("Perimeter\ngot: %.2f\nexpect: %.2f", got, want)
		}
	})

}

type TableAreaTest struct {
	geoForm Geoform
	want    float64
}

func TestArea(t *testing.T) {
	testArea := []TableAreaTest{
		{geoForm: &Retangule{12, 6}, want: 72.0},
		{geoForm: &Circule{10}, want: 314.1592653589793},
		{geoForm: &Triangule{12, 6}, want: 36.0},
	}

	for _, tt := range testArea {
		got := tt.geoForm.Area()

		if got != tt.want {
			t.Errorf("Perimeter of %#v\ngot: %.2f\nexpect: %.2f", tt.geoForm, got, tt.want)
		}
	}
}
