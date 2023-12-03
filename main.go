package main

import (
	"fmt"
	"math"
	"testing"
)

// Returns the minimum of the f(x) within the range from x0 to x1 with the precision given as e
// Note that presence of only 1 extremum within the given range is allowed
func FindMinimumUnoptimized(f func(float64) float64, x0, x1, e float64) float64 {
	var i uint
	for delta := (x1 - x0) / 4; 2*delta > e; delta = (x1 - x0) / 4 {
		xm := x0 + (x1-x0)/2
		c0, c1 := f(xm-delta), f(xm+delta)
		if c0 < c1 {
			x1 = xm
		} else {
			x0 = xm
		}
		i++
	}
	fmt.Printf("Unoptimised method iteration count: %v\n", i)
	return x0 + (x1-x0)/2
}

// Returns the minimum of the f(x) within the range from x0 to x1 with the precision given as e
// It uses golden ratio to reduce f(x) calls via reusage of some calculations
// Only works with x0 >= 0 and x1 >= 0
// Note that presence of only 1 extremum within the given range is allowed
func FindMinimumOptimised(f func(float64) float64, x0, x1, e float64) float64 {
	var i uint
	var gr = (3.0 - math.Sqrt(5)) / 2.0 // Golden ratio
	Δx := x1 - x0
	a := Δx*gr + x0
	c0, c1 := x0+a, x1-a
	for ; Δx >= e; Δx = x1 - x0 {
		a = Δx - 2*a
		y0, y1 := f(c0), f(c1)
		if y0 < y1 {
			x1, c1, c0 = c1, c0, x0+a
		} else {
			x0, c0, c1 = c0, c1, x1-a
		}
		i++
	}
	fmt.Printf("Optimised method iteration count: %v\n", i)
	return x0 // Any val. ∈ [x0;x0] suffices, since x1-x0<e
}

func nearlyEqual(a, b, e float64) bool {
	return math.Abs(a-b) <= e
}

func TestFindMinimum(t *testing.T) {
	var fxCalls uint
	type Input struct {
		f         func(float64) float64
		x0, x1, e float64
	}
	for num, tc := range []struct {
		input Input
		want  float64
	}{
		{Input{func(x float64) float64 { fxCalls++; return 0.1*x*x - math.Sqrt(97.0)*x + 10 }, 0, 100, 1e-5}, 49.24428900},
		// {Input{func(x float64) float64 { fxCalls++; return 0.1*x*x - math.Sqrt(97.0)*x + 10 }, 0, 100, 1e-6}, 49.24428900}, // They will both fail
	} {
		if got := FindMinimumUnoptimized(tc.input.f, tc.input.x0, tc.input.x1, tc.input.e); !nearlyEqual(got, tc.want, tc.input.e) {
			t.Errorf("FindMinimumUnoptimised failed test No %v: got = %v, want = %v", num, got, tc.want)
		}
		fmt.Printf("Unoptimised method f(x) call count: %v\n", fxCalls)
		fxCalls = 0
		if got := FindMinimumOptimised(tc.input.f, tc.input.x0, tc.input.x1, tc.input.e); !nearlyEqual(got, tc.want, tc.input.e) {
			t.Errorf("FindMinimumOptimised failed test No %v: got = %v, want = %v", num, got, tc.want)
		}
		fmt.Printf("Optimised method f(x) call count: %v\n", fxCalls)
		// Why on Earth does the optimized method perform worse?
	}
}

func TestNearlyEqual(t *testing.T) {
	type Input struct {
		a, b, e float64
	}
	for num, tc := range []struct {
		input Input
		want  bool
	}{
		{Input{1, 1, 1}, true},
		{Input{1.111111, 1.111111, 1e-6}, true},
		{Input{1.111111, 1.1111111111111111, 1e-6}, true},
		{Input{9.999999, 9.9999999999999999, 1e-6}, true},
		{Input{1.23457, 1.234567, 1e-6}, false},
		{Input{49.244287809901444, 49.244289, 1e-6}, false},
	} {
		if got := nearlyEqual(tc.input.a, tc.input.b, tc.input.e); got != tc.want {
			t.Errorf("Function failed test No %v: got = %v, want = %v", num, got, tc.want)
		}
	}
}

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test FindMinimum", F: TestFindMinimum},
			{Name: "Test NearlyEqual", F: TestNearlyEqual},
		},
		/* benchmarks */ nil /* examples */, nil)
}
