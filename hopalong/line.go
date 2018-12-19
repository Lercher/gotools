package main

import "math"

type line struct {
	min, max float64
}

func (l *line) enclose(val float64) {
	l.min = math.Min(l.min, val)
	l.max = math.Max(l.max, val)
}

func (l *line) linetransform(w int) *linetransform {
	// create a linear transform that maps min to 0 and max to w
	return &linetransform{
		offset: -l.min,
		factor: float64(w) / (l.max - l.min),
	}
}