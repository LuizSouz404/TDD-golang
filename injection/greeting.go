package main

import (
	"fmt"
	"io"
)

func main() {}

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Olá, %s", name)
}
