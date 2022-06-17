package main

import (
    "testing"
    "bytes"
    "reflect"
    "time"
)

func TestCountdown(t *testing.T) {
    t.Run("Number of calls", func (t *testing.T) {
        buffer := &bytes.Buffer{}
        spy := &SpyCountdownOperations{}

        Countdown(buffer, spy)

        got := buffer.String()
        want := "3\n2\n1\nGo!"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
    t.Run("Number of calls", func (t *testing.T) {
        spy := &SpyCountdownOperations{}

        Countdown(spy, spy)

        want := []string{
            "write",
            "sleep",
            "write",
            "sleep",
            "write",
            "sleep",
            "write",
        }

        if !reflect.DeepEqual(want, spy.Calls) {
		    t.Errorf("wanted calls %v got %v", want, spy.Calls)
        }

    })
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
