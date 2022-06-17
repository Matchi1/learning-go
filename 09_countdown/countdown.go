package main

import (
    "io"
    "os"
    "fmt"
    "time"
)

const finalWord = "Go!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
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
