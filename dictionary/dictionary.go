package main

import "errors"

func main() {}

type Dictionary map[string]string

var (
	ErrorWordNotFound = errors.New("word not found")
	ErrorWordExists   = errors.New("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	if d[word] == "" {
		return "", ErrorWordNotFound
	}
	return d[word], nil
}

func (d Dictionary) Add(word string, definition string) error {
	if d[word] == definition {
		return ErrorWordExists
	}

	d[word] = definition
	return nil
}
