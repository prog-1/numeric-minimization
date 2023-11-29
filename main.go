package main

import (
	"fmt"
	"math"
)

func s(x float64) float64 {
	return 0.1*x*x - math.Sqrt(97)*x + 10
}

func FindMinBisect(f func(x float64) float64, startx, endx float64, e float64) float64 {
	d := e / 2
	for (endx - startx) > e {
		p := (startx + endx) / 2
		if f(p-d) < f(p+d) {
			endx = p
		} else {
			startx = p
		}
	}
	return startx
}

func main() {
	fmt.Println(FindMinBisect(s, 20, 80, 0.000001))
}
