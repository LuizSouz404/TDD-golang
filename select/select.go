package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {}

var limitTenSecond = 10 * time.Second

func Runner(urlOne, urlTwo string) (winner string, erro error) {
	return ConfigRunner(urlOne, urlTwo, limitTenSecond)
}

func ConfigRunner(urlOne, urlTwo string, limitTime time.Duration) (winner string, erro error) {
	select {
	case <-mensureTimeResponse(urlOne):
		return urlOne, nil
	case <-mensureTimeResponse(urlTwo):
		return urlTwo, nil
	case <-time.After(limitTime):
		return "", fmt.Errorf("limit time of waiting was exceeded")
	}
}

func mensureTimeResponse(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}
