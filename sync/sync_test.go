package main

import "testing"

func TestSync(t *testing.T) {
	t.Run("Increment counter 3 times and result value should be 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		if counter.Value() != 3 {
			t.Errorf("Counter\ngot: %d\nexpect: %d", counter.Value(), 3)
		}
	})
}
