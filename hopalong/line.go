package main

import "math"

type line struct {
	min, max float64
}

func (l *line) enclose(val float64) {
	l.min = math.Min(l.min, val)
	l.max = math.Max(l.max, val)
}
