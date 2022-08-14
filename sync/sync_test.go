package main

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	t.Run("Increment counter 3 times and result value should be 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		verifyCounter(t, counter, 3)
	})

	t.Run("should run concurrent in safety", func(t *testing.T) {
		counterExpect := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(counterExpect)

		for i := 0; i < counterExpect; i++ {
			go func(w *sync.WaitGroup) {
				counter.Increment()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verifyCounter(t, counter, counterExpect)
	})
}

func verifyCounter(t *testing.T, counter Counter, times int) {
	t.Helper()
	if counter.Value() != times {
		t.Errorf("Counter\ngot: %d\nexpect: %d", counter.Value(), times)
	}
}
