package main

import (
	"math"
	"testing"
)

func t1(x float64) float64 {
	return 0.1*x*x - math.Sqrt(97)*x + 10
}

func t2(x float64) float64 {
	return x*x - 10*x
}
func t3(x float64) float64 {
	return x*x + 10*x
}

func TestFindMinBisect(t *testing.T) {
	for _, tc := range []struct {
		name         string
		f            func(x float64) float64
		startx, endx float64
		eps          float64
		res          float64
	}{
		{"0.1x^2-sqrt(97)*x+10", t1, 20, 80, 1e-5, 49.244289008980523608731057074588122408480681437213820858615772649182202918538393150466003920577128812190505994358745998590375518690734751424644613904864435},
		{"x^2 - 10x", t2, 0, 20, 1e-5, 5},
		{"x^2 + 10x", t3, -20, 20, 1e-5, -5},
	} {
		if res := FindMinBisect(tc.f, tc.startx, tc.endx, tc.eps); math.Abs(tc.res-res) > tc.eps {
			t.Errorf("FindMinBisect(%s, %f, %f, %f) = %f, want ~ %f`", tc.name, tc.startx, tc.endx, tc.eps, res, tc.res)
		}
	}
}
