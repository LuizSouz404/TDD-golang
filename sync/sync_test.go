package main

import "testing"

func TestSync(t *testing.T) {
	t.Run("Increment counter 3 times and result value should be 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		verifyCounter(t, counter, 3)
	})
}

func verifyCounter(t *testing.T, counter Counter, times int) {
	if counter.Value() != times {
		t.Errorf("Counter\ngot: %d\nexpect: %d", counter.Value(), 3)
	}
}
