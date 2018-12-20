package main

import "math"

type linetransform struct {
	factor, offset float64
	max            int
}

func (t *linetransform) tr(val float64) (int, bool) {
	i := int(math.Floor((val + t.offset) * t.factor))
	return i, 0<=i && i<t.max
}
