package main

// go get && go build && ./hopalong -w 300 -h 300 -o x.png && rm hopalong

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
	flagFile      = flag.String("o", "", "output the png image to this `file`, if missing or -, stdout is used")
)

func main() {
	log.Println("This is HopAlong, (c) 2018 by Martin Lercher")
	flag.Parse()

	t := time.Now()
	defer func() {
		log.Println("runtime:", time.Now().Sub(t))
	}()

	o := os.Stdout
	if *flagFile != "" && *flagFile != "-" {
		f, err := os.Create(*flagFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		o = f
	}
	hopPNG(o, *flagWidth, *flagHeight, *flagRounds, *flagNextColor)
}
