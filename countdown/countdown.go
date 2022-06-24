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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *ConfigurableSleeper) Sleep() {
    s.sleep(s.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countDownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
        sleeper.Sleep()
    }
    fmt.Fprint(out, finalWord)
}

func main() {
    sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}
