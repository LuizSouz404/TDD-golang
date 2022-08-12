package main

import (
	"net/http"
	"time"
)

func main() {}

func Runner(urlOne, urlTwo string) string {
	initA := time.Now()
	http.Get(urlOne)
	durationA := time.Since(initA)

	initB := time.Now()
	http.Get(urlTwo)
	durationB := time.Since(initB)

	if durationA < durationB {
		return urlOne
	}

	return urlTwo
}
