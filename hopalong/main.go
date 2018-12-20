package main

// go build && ./hopalong >x.png && rm hopalong

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	flagWidth     = flag.Int("w", 1920, "width of the png image")
	flagHeight    = flag.Int("h", 1920, "height of the png image")
	flagRounds    = flag.Int("n", 5000000, "number of plotted orbit pixels")
	flagNextColor = flag.Int("c", 1000, "change color every this number of iterations")
)

func main() {
	log.Println("This is HopAlong, (c) 2018 by Martin Lercher")
	flag.Parse()

	t := time.Now()
	defer func() {
		log.Println("runtime:", time.Now().Sub(t))
	}()

	hopPNG(os.Stdout, *flagWidth, *flagHeight, *flagRounds, *flagNextColor)
}
