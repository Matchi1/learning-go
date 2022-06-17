package main

import (
    "io"
    "os"
    "fmt"
    "time"
)

const finalWord = "Go!"
const countDownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countDownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
        sleeper.Sleep()
    }
    fmt.Fprint(out, finalWord)
}

func main() {
    sleeper := &DefaultSleeper{}
    Countdown(os.Stdout, sleeper)
}
