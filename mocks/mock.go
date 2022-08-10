package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &SleeperDefault{}
	Count(os.Stdout, sleeper)
}

const lastWord = "Vai!"
const countNumber = 3

type Sleeper interface {
	Sleep()
}

type SleeperConfigure struct {
	duration time.Duration
	pause    func(time.Duration)
}

func (s *SleeperConfigure) Pause() {
	s.pause(s.duration)
}

type SleeperDefault struct{}

func (s *SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

func Count(writer io.Writer, sleeper Sleeper) {
	for i := countNumber; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, lastWord)
}
