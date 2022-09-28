package main

import "testing"

func TestSumdigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{333, 9},
		{321, 6},
		{21, 3},
		{10000, 1},
		{420, 6},
		{228, 12},
	} {
		if got := sumdigits(tc.n); got != tc.want {
			t.Errorf("numSum(%v)   got = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
