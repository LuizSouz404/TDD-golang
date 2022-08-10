package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	buffer := bytes.Buffer{}
	Count(&buffer)

	got := buffer.String()
	want := `3
2
1
Vai!`

	if got != want {
		t.Errorf("Count\ngot: %s\nexpect: %s", got, want)
	}
}
