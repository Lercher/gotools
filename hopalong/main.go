package main

// go get && go build && ./hopalong -w 300 -h 300 -f 6 -q1 -o x.png && rm hopalong

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	flagWidth     = flag.Int("w", 1920, "width of the png image")
	flagFunc      = flag.Int("f", 0, fmt.Sprintf("use this hop function in the range 0..%d", len(hopfuncs)))
	flagHeight    = flag.Int("h", 1920, "height of the png image")
	flagRounds    = flag.Int("n", 5000000, "number of plotted orbit pixels")
	flagNextColor = flag.Int("c", 1000, "change color every this number of iterations")
	flagFile      = flag.String("o", "", "output the png image to this `file`, if missing or -, stdout is used")
	flagQ1        = flag.Bool("q1", false, "render only the top left quadrant pixels")
)

func main() {
	log.Println("This is HopAlong, (c) 2018 by Martin Lercher")
	flag.Parse()

	if *flagFunc <0 || len(hopfuncs) <= *flagFunc {
		log.Fatalln("-f", *flagFunc, "out of range")
	}

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
	h := hopPNG(o,*flagFunc, *flagWidth, *flagHeight, *flagRounds, *flagNextColor, *flagQ1)
	log.Printf("A%v B%v C%v D%v, box: x%v y%v", h.a, h.b, h.c, h.d, h.box.x, h.box.y)
}
