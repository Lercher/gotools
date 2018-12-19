package main

import "math"

type hop struct {
	a, b, c float64 // constant parameters
	x, y    float64 // iteration coordinates
	box             // initial bounding box is only the origin at 0,0
}

func (h *hop) configure(a, b, c float64) {
	h.a, h.b, h.c = a, b, c
}

func (h *hop) rounds(n int) {
	for i := 0; i < n; i++ {
		s := math.Sqrt(math.Abs(h.b*h.x - h.c))
		if h.x >= 0 {
			s = -s
		}
		h.x, h.y = h.y+s, h.a-h.x
		h.enclose(h.x, h.y) // enlarge bounding box
	}
}

func (h *hop) randomizeABC() {
	h.a = ran(40.0, 1540.0)
	h.b = ran(3.0, 20.0)
	h.c = ran(100.0, 3100.0)
}
