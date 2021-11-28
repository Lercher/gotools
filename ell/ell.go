package main

import (
	"fmt"
	"log"
)

const eps = 5e-2

func main() {
	log.Println("This is ell, (c) 2021 by Martin Lercher")

	printwalk(3, 120)
}

func printwalk(w float64, steps int) {
	y0 := float64(0)
	walk00(w, steps, func(_, y float64, isZero bool) {
		if y0 != y {
			y0 = y
			fmt.Println()
		}
		if isZero {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	})
}

func walk00(w float64, steps int, result func(x, y float64, isZero bool)) {
	delta := w / float64(steps)
	for j := 0; j < steps; j++ {
		y := -w/2 + float64(j)*delta
		for i := 0; i < steps; i++ {
			x := -w/2 + float64(i)*delta
			result(x, y, isZeroAt(x, y))
		}
	}
}

func isZeroAt(x, y float64) bool {
	v := f0(x, y)
	return -eps <= v && v <= eps
}

func f0(x, y float64) float64 {
	return -y*y + x*x*x - x + 0.45
}
