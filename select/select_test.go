package main

import "testing"

func TestSelect(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://quii.com.uk"

	want := fastURL
	got := Runner(slowURL, fastURL)

	if got != want {
		t.Errorf("Runner\ngot: %s\nexpect: %s", got, want)
	}
}
