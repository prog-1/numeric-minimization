package main

import (
	"fmt"
	"math"
)

var (
	goldenRatio = (3 - math.Sqrt(5)) / 2
)

func t1(x float64) float64 {
	return 0.1*x*x - math.Sqrt(97)*x + 10
}

func FindMinGoldenRatio(f func(x float64) float64, startx, endx float64, e float64) float64 {
	left, right := startx+(endx-startx)*goldenRatio, endx-(endx-startx)*goldenRatio
	leftv, rightv := f(left), f(right)
	for (endx - startx) > e {
		if leftv < rightv {
			endx = right
			right = left
			rightv = leftv
			left = startx + (endx-startx)*goldenRatio
			leftv = f(left)
		} else {
			startx = left
			left = right
			leftv = rightv

			right = endx - (endx-startx)*goldenRatio
			rightv = f(right)
		}
	}
	return startx
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
	fmt.Println(FindMinGoldenRatio(t1, 20, 80, 0.000001))
	fmt.Println(FindMinBisect(t1, 20, 80, 0.000001))

}
