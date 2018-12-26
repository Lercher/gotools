package main

import (
	"math"
)

type hopfunc func(h *hop)

var hopfuncs = []hopfunc{
	martin1,
	martin2,
	ejk1,
	ejk2,
	ejk3,
	ejk4,
	ejk5,
	ejk6,
	rr1,
}

var randfuncs = []hopfunc{
	rmartin1,
	rmartin2,
	rejk1,
	rejk2,
	rejk3,
	rejk4,
	rejk5,
	rejk6,
	rrr1,
}

func rmartin1(h *hop) {
	h.a = ran(40, 1540)
	h.b = ran(3, 20)
	h.c = ran(100, 3100)
}

func rmartin2(h *hop) {
	h.a = ran(3.0715927, 3.2115927)
}

func rejk1(h *hop) {
	h.a = ran(0, 500)
	h.b = ran(0, 0.4)
	h.c = ran(10, 110)
}

func rejk2(h *hop) {
	h.a = ran(0, 500)
	h.b = math.Pow(10, 6+ran(0, 24))
	if rng.Float32() < 0.5 {
		h.b = -h.b
	}
	h.c = math.Pow(10, ran(0, 9))
	if rng.Float32() < 0.5 {
		h.c = -h.c
	}
}

func rejk3(h *hop) {
	h.a = ran(0, 500)
	h.b = ran(0.05, 0.40)
	h.c = ran(30, 110)
}

func rejk4(h *hop) {
	h.a = ran(0, 1000)
	h.b = ran(1, 10)
	h.c = ran(30, 70)
}

func rejk5(h *hop) {
	h.a = ran(0, 600)
	h.b = ran(0.1, 0.4)
	h.c = ran(20, 110)
}

func rejk6(h *hop) {
	h.a = ran(550, 650)
	h.b = ran(0.5, 1.5)
}

func martin1(h *hop) {
	s := math.Sqrt(math.Abs(h.b*h.x - h.c))
	if h.x >= 0 {
		s = -s
	}
	h.x, h.y = h.y+s, h.a-h.x

}

func martin2(h *hop) {
	h.x, h.y = h.y-math.Sin(h.x), h.a-h.x
}

func ejk1(h *hop) {
	s := h.b*h.x - h.c
	if h.x >= 0 {
		s = -s
	}
	h.x, h.y = h.y+s, h.a-h.x
}

func ejk2(h *hop) {
	s := math.Log(math.Abs(h.b*h.x - h.c))
	if h.x >= 0 {
		s = -s
	}
	h.x, h.y = h.y-s, h.a-h.x
}

func ejk3(h *hop) {
	s := math.Sin(h.b*h.x) - h.c
	if h.x >= 0 {
		s = -s
	}
	h.x, h.y = h.y+s, h.a-h.x
}

func ejk4(h *hop) {
	var s float64
	if h.x >= 0 {
		s = math.Sin(h.b*h.x) - h.c
	} else {
		s = -math.Sqrt(math.Abs(h.b*h.x - h.c))
	}
	h.x, h.y = h.y-s, h.a-h.x
}

func ejk5(h *hop) {
	var s float64
	if h.x >= 0 {
		s = math.Sin(h.b*h.x) - h.c
	} else {
		s = -(h.b*h.x - h.c)
	}
	h.x, h.y = h.y-s, h.a-h.x
}

func ejk6(h *hop) {
	s := h.b * h.x
	h.x, h.y = h.y-math.Asin(s-math.Floor(s)), h.a-h.x
}

func rrr1(h *hop) {
	h.a = ran(0, 100.0)
	h.b = ran(0, 20.0)
	h.c = ran(0, 200.0)
	h.d = ran(0.0, 0.9) // used as power exponent in rr1 only
}

func rr1(h *hop) {
	s := math.Pow(math.Abs(h.b*h.x-h.c), h.d)
	if h.x >= 0 {
		s = -s
	}
	h.x, h.y = h.y-s, h.a-h.x
}
