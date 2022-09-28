package main

import "testing"

func TestIspowerof2(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{512, true},
		{128, true},
		{9, false},
		{3, false},
		{64, true},
		{31, false},
	} {
		if got := ispowerof2(tc.n); got != tc.want {
			t.Errorf("powerof2(%v)   got = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
