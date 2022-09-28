package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		res1 int
		res2 int
	}{
		{[]int{20, 40, 20, 11, 38}, 40, 38},
		{[]int{10, 11, 13, 1, 2}, 13, 11},
		{[]int{40, 40, 40}, 40, 40},
		{[]int{10, 11, 13, 1, 2}, 13, 11},
		{[]int{40, 10, 20, 20}, 40, 20},
		{[]int{10, 11, 13, 1, 2}, 13, 11},
		{[]int{40, 1}, 40, 1},
		{[]int{10, 11, 13, 1, 2}, 13, 11},
		{[]int{10, 11, 13, 1, 2}, 13, 11},
		{[]int{-4, 2, 3}, 3, 2},
	} {
		if max1, max2 := max(tc.s); max1 != tc.res1 || max2 != tc.res2 {
			t.Errorf("got = %v, %v, want = %v, %v", max1, max2, tc.res1, tc.res2)
		}
	}
}
