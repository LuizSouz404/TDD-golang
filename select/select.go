package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {}

func Runner(urlOne, urlTwo string) (winner string, erro error) {
	select {
	case <-mensureTimeResponse(urlOne):
		return urlOne, nil
	case <-mensureTimeResponse(urlTwo):
		return urlTwo, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("Limit time of waiting was exceeded")
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
