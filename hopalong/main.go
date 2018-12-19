package main

import (
	"log"
	"time"
)

func main() {
	log.Println("This is HopAlong, (c) 2018 by Martin Lercher")

	t := time.Now()
	defer func() {
		log.Println("runtime:", time.Now().Sub(t))
	}()
	h := &hop{}
	h.randomizeABC()
	h.rounds(1000000)
	log.Println(*h)
}
