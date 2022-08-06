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

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Geoform, want float64) {
		t.Helper()
		got := forma.Area()

		if got != want {
			t.Errorf("Perimeter\ngot: %.2f\nexpect: %.2f", got, want)
		}
	}

	t.Run("this should returns area of triangule", func(t *testing.T) {
		retangule := Retangule{12.0, 6.0}
		want := 72.0

		verificaArea(t, &retangule, want)
	})

	t.Run("this shoul reurn area of circule", func(t *testing.T) {
		circule := Circule{10.0}
		want := 314.1592653589793

		verificaArea(t, &circule, want)
	})
}
