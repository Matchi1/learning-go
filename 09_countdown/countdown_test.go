package main

import (
    "testing"
    "bytes"
    "reflect"
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
