package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}
	Count(&buffer, sleeperSpy)

	got := buffer.String()
	want := `3
2
1
Vai!`

	if got != want {
		t.Errorf("Count times: %d\ngot: %s\nexpect: %s", sleeperSpy.Duration, got, want)
	}
}
