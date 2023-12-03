package main

import (
	"math"
	"testing"
)

func TestFindExtremum(t *testing.T) {
	for _, tc := range []struct {
		a, b, eps float64
		want      float64
	}{
		{a: 0, b: 10, eps: 0.0001, want: 9.9999},
		{a: 45, b: 55, eps: 0.0001, want: 49.2443},
		{a: 100, b: 155, eps: 0.1, want: 100.1},
	} {
		got := findExtremum(tc.a, tc.b, tc.eps)
		if math.Abs(got-tc.want) > tc.eps {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}

	}
}

func TestFindExtremum2(t *testing.T) {
	for _, tc := range []struct {
		a, b, eps float64
		want      float64
	}{
		{a: 0, b: 10, eps: 0.0001, want: 9.9999},
		{a: 45, b: 55, eps: 0.0001, want: 49.2443},
		{a: 100, b: 155, eps: 0.1, want: 100.1},
	} {
		got := findExtremum2(tc.a, tc.b, tc.eps)
		if math.Abs(got-tc.want) > tc.eps {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}

	}
}
