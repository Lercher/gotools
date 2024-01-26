package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This program calculates some E(x)")

	n := 100000
	k := 0
	sum := 0
	m := make(map[int]int)
	for n > 0 {
		n--

		i := 0
		for {
			i++
			r := rand.Float32()
			if r < 0.5 {
				break
			}
		}
		m[i]++
		k++
		sum += i
	}
	fmt.Println("E(x)=", float64(sum)/float64(k), "for", k, "cycles. Histogram:", m)
}
