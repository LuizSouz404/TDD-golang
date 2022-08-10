package main

import (
	"bytes"
	"reflect"
	"testing"
)

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
