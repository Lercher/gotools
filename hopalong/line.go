package main

import "math"

type line struct {
	min, max float64
}

func (l *line) enclose(val float64) {
	l.min = math.Min(l.min, val)
	l.max = math.Max(l.max, val)
}

func (l *line) linetransform(w int, quadrant1 bool) *linetransform {
	// quadrant1=true;  create a linear transform that maps 0 to 0 and max to w
	if quadrant1 {
		return &linetransform{
			offset: 0,
			factor: float64(w) / l.max,
			max:    w,
		}
	}
	// quadrant1=false; create a linear transform that maps min to 0 and max to w
	return &linetransform{
		offset: -l.min,
		factor: float64(w) / (l.max - l.min),
		max:    w,
	}
}
