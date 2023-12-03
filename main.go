package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 0.1*x*x - math.Sqrt(97)*x + 10
}

func findExtremum(a, b, e float64) float64 {
	iterations := 0
	for math.Abs(b-a) > e {
		iterations++
		c := (a + b) / 2
		fa, fc := f(a), f(c)
		if fc > fa {
			b = c
		} else {
			a = c
		}
	}
	fmt.Println(iterations)
	return (a + b) / 2
}

// golden ratio method
func findExtremum2(a, b, e float64) float64 {
	iterations := 0
	g := (1 + math.Sqrt(5)) / 2
	x1, x2 := b-(b-a)/g, a+(b-a)/g
	f1, f2 := f(x1), f(x2)
	for math.Abs(b-a) > e {
		if f1 >= f2 {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + (b-a)/g
			f2 = f(x2)
		} else {
			b = x2
			x2 = x1
			f2 = f1
			x1 = b - (b-a)/g
			f1 = f(x1)
		}
		iterations++
	}
	fmt.Println(iterations)
	return (a + b) / 2
}

func main() {
	a := 0.0
	b := 60.0
	eps := 0.0001
	fmt.Println(findExtremum(a, b, eps))
	fmt.Println(findExtremum2(a, b, eps))
}
