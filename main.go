package main

import (
	"math"
	"testing"
)

// f(x) = 0.1x^2-sqrt(97)*x+10
// f'(x) = 0.2x - sqrt(97)

// TODO: Compare 2 methods

// Returns the minimum of the f(x) within the range from x0 to x1 using f(x) derivative f_(x) with the precision given as e
func FindMinimum(f, f_ func(float64) float64, x0, x1, e float64) float64 {
	e /= 10
	for delta := (x1 - x0) / 4; 2*delta > e; delta = (x1 - x0) / 4 {
		xm := x0 + (x1-x0)/2
		c0, c1 := f(xm-delta), f(xm+delta)
		if c0 < c1 {
			x1 = xm
		} else {
			x0 = xm
		}
	}
	return x0 + (x1-x0)/2
}

func nearlyEqual(a, b, e float64) bool {
	return math.Abs(a-b) <= e
}

func TestFindMinimum(t *testing.T) {
	type Input struct {
		f, f_     func(float64) float64
		x0, x1, e float64
	}
	sqrt97 := math.Sqrt(97.0)
	for num, tc := range []struct {
		input Input
		want  float64
	}{
		{Input{func(x float64) float64 { return 0.1*x*x - sqrt97*x + 10 }, func(x float64) float64 { return 0.2*x - sqrt97 }, 0, 100, 1e-7}, 49.2442900},
	} {
		if got := FindMinimum(tc.input.f, tc.input.f_, tc.input.x0, tc.input.x1, tc.input.e); !nearlyEqual(got, tc.want, tc.input.e) {
			t.Errorf("FindMinimum failed test No %v: got = %v, want = %v", num, got, tc.want)
		}
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
		// {Input{1, 1, 1}, true},
		// {Input{1.111111, 1.111111, 1e-6}, true},
		// {Input{1.111111, 1.1111111111111111, 1e-6}, true},
		// {Input{9.999999, 9.9999999999999999, 1e-6}, true},
		{Input{1.23456, 1.234567, 1e-6}, true},
		// {Input{1.23457, 1.234567, 1e-6}, false},
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
