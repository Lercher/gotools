package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("This is ell, (c) 2021 by Martin Lercher")

	// printwalk(3, 120, 5e-2)
	o, err := os.Create("img.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer o.Close()

	ellPNG(o, 3, 1920, 3e-2)
	log.Println("done")
}

func printwalk(w float64, steps int, epsilon float64) {
	j0 := 0
	walk00(w, steps, epsilon, func(_, j int, _, _, _ float64, isZero bool) {
		if j0 != j {
			j0 = j
			fmt.Println()
		}
		if isZero {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	})
}

func walk00(w float64, steps int, epsilon float64, result func(i, j int, x, y, val float64, isZero bool)) {
	delta := w / float64(steps)
	for j := 0; j < steps; j++ {
		y := -w/2 + float64(j)*delta
		for i := 0; i < steps; i++ {
			x := -w/2 + float64(i)*delta
			v, zero := isZeroAt(x, y, epsilon)
			result(i, j, x, y, v, zero)
		}
	}
}

func isZeroAt(x, y, epsilon float64) (float64, bool) {
	v := f0(x, y)
	return v, -epsilon <= v && v <= epsilon
}

func f0(x, y float64) float64 {
	return -y*y + x*x*x - x + 0.42
}
