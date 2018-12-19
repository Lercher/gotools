package main

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func ran(a, b float64) float64 {
	return a + (b-a)*rng.Float64()
}
