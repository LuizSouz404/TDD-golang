package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	t.Run("This should test if implemntations count is rigth", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}
		Count(&buffer, sleeperSpy)

		got := buffer.String()
		want := `3
2
1
Vai!`

		if got != want {
			t.Errorf("Count times: %d\ngot: %s\nexpect: %s", sleeperSpy.Times, got, want)
		}
	})

	t.Run("This should to guarantee the sequence is rigth", func(t *testing.T) {
		spyImpressSleep := &SpyCountOperations{}
		Count(spyImpressSleep, spyImpressSleep)

		want := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}

		if !reflect.DeepEqual(want, spyImpressSleep.Times) {
			t.Errorf("Count\ngot: %v\nexpect: %v", spyImpressSleep.Times, want)
		}
	})
}

func TestSleeperConfigure(t *testing.T) {
	timePause := 5 * time.Second

	timeSpy := &TimeSpy{}
	sleeper := SleeperConfigure{timePause, timeSpy.Pause}
	sleeper.Pause()

	if timeSpy.durationPause != timePause {
		t.Errorf("Should pause for %v, but pause for %v", timePause, timeSpy.durationPause)
	}
}

const pause = "pause"
const write = "write"

type SleeperSpy struct {
	Times int
}

func (s *SleeperSpy) Sleep() {
	s.Times++
}

type SpyCountOperations struct {
	Times []string
}

func (s *SpyCountOperations) Sleep() {
	s.Times = append(s.Times, pause)
}

func (s *SpyCountOperations) Write(p []byte) (n int, err error) {
	s.Times = append(s.Times, write)
	return
}

type TimeSpy struct {
	durationPause time.Duration
}

func (t *TimeSpy) Pause(duration time.Duration) {
	t.durationPause = duration
}
