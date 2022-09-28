package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		wantMax1 int
		wantMax2 int
	}{
		{[]int{0, 0, 0}, 0, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{-1, -2, -3, -4, -5}, -1, -2},
	} {
		if max1, max2 := Max(tc.input); max1 != tc.wantMax1 || max2 != tc.wantMax2 {
			t.Errorf("got = %v, %v, want = %v, %v", max1, max2, tc.wantMax1, tc.wantMax2)
		}
	}
}
