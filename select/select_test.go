package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	t.Run("this test should return an url", func(t *testing.T) {
		slowServer := createServerWithDelay(20 * time.Millisecond)
		fastServer := createServerWithDelay(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got, _ := Runner(slowURL, fastURL)

		if got != want {
			t.Errorf("Runner\ngot: %s\nexpect: %s", got, want)
		}
	})

	t.Run("return an erro if an server not response in 10s", func(t *testing.T) {
		serverA := createServerWithDelay(11 * time.Second)
		serverB := createServerWithDelay(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Runner(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expect an error, but nothing was got")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
