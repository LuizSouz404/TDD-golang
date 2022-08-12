package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	slowServer := createServerWithDelay(20 * time.Millisecond)
	fastServer := createServerWithDelay(0 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	defer slowServer.Close()
	defer fastServer.Close()

	want := fastURL
	got := Runner(slowURL, fastURL)

	if got != want {
		t.Errorf("Runner\ngot: %s\nexpect: %s", got, want)
	}
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
