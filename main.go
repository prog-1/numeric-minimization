package main

import (
	"fmt"
)

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
	fmt.Println(FindMinBisect(t1, 20, 80, 0.000001))
}
