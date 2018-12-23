package main

type hop struct {
	hopf       hopfunc
	a, b, c, d float64 // constant parameters
	x, y       float64 // iteration coordinates
	box                // initial bounding box is only the origin at 0,0
}

func (h *hop) reset() {
	h.x, h.y = 0, 0
	h.rounds(10, func() {}) // ignore some first points of the hop function's orbit
}

func (h *hop) bounds(n int) {
	h.reset()
	h.rounds(n, func() {
		h.enclose(h.x, h.y) // enlarge bounding box
	})
}

func (h *hop) rounds(n int, f func()) {
	for i := 0; i < n; i++ {
		h.hopf(h)
		f()
	}
}
