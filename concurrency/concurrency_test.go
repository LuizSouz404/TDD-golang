package main

import (
	"reflect"
	"testing"
)

func mockCheckerWebsite(url string) bool {
	if url == "waat://furhuterwe.geds" {
		return false
	}

	return true
}

func TestVerifyWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhuterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhuterwe.geds":     false,
	}

	got := VerifyWebsite(mockCheckerWebsite, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Concurrency\ngot: %v\nwant: %v", got, want)
	}
}
