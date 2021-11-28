package main

import "testing"

var r float64

func BenchmarkElliptic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := float64(i) / float64(b.N)
		y := 1 - x
		r = f0(x, y)
	}
}
