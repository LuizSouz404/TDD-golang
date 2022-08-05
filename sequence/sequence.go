package main

func main() {}

func Sequence(value string, times int) string {
	sequence := ""

	if value == "" {
		value = "a"
	}

	for i := 0; i < times; i++ {
		sequence += value
	}
	return sequence
}
