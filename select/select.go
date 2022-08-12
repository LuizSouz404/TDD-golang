package main

import (
	"net/http"
	"time"
)

func main() {}

func Runner(urlOne, urlTwo string) string {
	durationA := mensureTimeResponse(urlOne)
	durationB := mensureTimeResponse(urlTwo)

	if durationA < durationB {
		return urlOne
	}

	return urlTwo
}

func mensureTimeResponse(URL string) time.Duration {
	init := time.Now()
	http.Get(URL)
	return time.Since(init)
}
