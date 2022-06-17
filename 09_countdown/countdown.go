package main

import (
    "io"
    "os"
    "fmt"
    "time"
)

const finalWord = "Go!"
const countDownStart = 3

func Countdown(out io.Writer) {
    for i := countDownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
        time.Sleep(1 * time.Second)
    }
    fmt.Fprint(out, finalWord)
}

func main() {
    Countdown(os.Stdout)
}
