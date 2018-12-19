package main

import "math"

type linetransform struct {
	factor, offset float64
}

func (t *linetransform) tr(val float64) int {
	return int(math.Round((val + t.offset) * t.factor))
}
