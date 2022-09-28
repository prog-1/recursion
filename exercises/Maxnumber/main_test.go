package main

import "testing"

func TestMaxnumber(t *testing.T) {
	for _, tc := range []struct {
		n    []int
		max  int
		i    int
		want int
	}{
		{[]int{1000, 1001, -100, 1}, 1000, 1, 1001},
		{[]int{1, -3, 99, 675}, 1, 1, 675},
		{[]int{10, 11, -100, 1}, 10, 1, 11},
	} {
		if max := maxnumber(tc.n, tc.max, tc.i); max != tc.want {
			t.Errorf("numSum(%v) got = %v, want = %v", tc.n, max, tc.want)
		}
	}
}

func TestMaxnumber2nd(t *testing.T) {
	for _, tc := range []struct {
		n    []int
		max  int
		max2 int
		i    int
		want int
	}{
		{[]int{1000, 1001, -100, 1}, 1000, 1001, 2, 1000},
		{[]int{1, -3, 99, 675}, 1, -3, 2, 99},
		{[]int{10, 11, -100, 1}, 10, 11, 2, 10},
	} {
		if max := maxnumber2nd(tc.n, tc.max, tc.max2, tc.i); max != tc.want {
			t.Errorf("numSum(%v) got = %v, want = %v", tc.n, max, tc.want)
		}
	}
}
