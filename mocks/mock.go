package main

import (
	"bytes"
	"fmt"
)

func main() {
}

func Count(writer *bytes.Buffer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, "Vai!")
}
