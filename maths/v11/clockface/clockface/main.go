package main

import (
	"os"
	"time"

	"github.com/Matchi1/learning-go/maths/v11/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
