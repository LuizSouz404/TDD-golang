package main

import (
	"bytes"
	"fmt"
)

func main() {
}

func Count(writer *bytes.Buffer) {
	fmt.Fprintf(writer, "3")
}
