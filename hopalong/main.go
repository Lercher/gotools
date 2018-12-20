package main

// go build && ./hopalong >x.png && rm hopalong

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Println("This is HopAlong, (c) 2018 by Martin Lercher")

	t := time.Now()
	defer func() {
		log.Println("runtime:", time.Now().Sub(t))
	}()

	hopPNG(os.Stdout, 1920, 1080, 100000)
}
