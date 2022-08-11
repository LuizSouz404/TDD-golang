package main

import (
	"reflect"
	"testing"
	"time"
)

func mockCheckerWebsite(url string) bool {
	return url != "waat://furhuterwe.geds"
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

func slowStubCheckerWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckerWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "one url"
	}

	for i := 0; i < b.N; i++ {
		VerifyWebsite(slowStubCheckerWebsite, urls)
	}
}
