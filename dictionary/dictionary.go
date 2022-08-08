package main

import "errors"

func main() {}

type Dictionary map[string]string

var (
	errorWordNotFound = errors.New("word not found")
)

func (d Dictionary) Search(word string) (string, error) {
	if d[word] == "" {
		return "", errorWordNotFound
	}
	return d[word], nil
}

func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}
