package main

import (
    "testing"
    "bytes"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
