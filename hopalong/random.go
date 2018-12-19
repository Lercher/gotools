package main

import "math/rand"

func ran(a, b float64) float64 {
	return a + (b-a) * rand.Float64()
}

