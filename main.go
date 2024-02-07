package main

import (
	"fmt"
	"log"
	"math"
)

func Bisection(a, b, eps float64, f func(float64) float64) float64 {
	if f(a)*f(b) >= 0 {
		log.Fatal("invalid interval")
	}
	t := (a + b) / 2 // the midpoint of a and b
	var count int
	for ; f(t) != 0 && math.Abs(b-a) > eps; t, count = (a+b)/2, count+1 {
		if f(t)*f(a) < 0 {
			b = t
		} else {
			a = t
		}
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*3+1) // because the program calculates f(t) 1 additional time before exiting the loop
	return t
}

func GoldenRatio(x0, x1, eps float64, f func(float64) float64) float64 {
	g := (1 + math.Sqrt(5)) / 2 // the golden ratio
	var count int
	for ; math.Abs(x1-x0) > eps; count++ {
		x2 := x0 + (g*(x1-x0))/2
		x3 := x1 - (g*(x1-x0))/2
		if f(x2) < f(x3) {
			x1 = x3
		} else {
			x0 = x2
		}
	}
	fmt.Printf("Number of iterations - %v, number of calculations f(x) - %v\n", count, count*2)
	return (x0 + x1) / 2
}

func main() {
	f := func(x float64) float64 { return 0.1*x*x - math.Sqrt(97)*x + 10 }
	fmt.Println(Bisection(-1, 5, 1e-3, f))
	fmt.Println(GoldenRatio(-1, 5, 1e-3, f))
}
