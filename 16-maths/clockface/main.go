package main

import (
	"os"
	"time"

	clockface "github.com/v1zix/learn-go-with-tests/16-maths/src"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
