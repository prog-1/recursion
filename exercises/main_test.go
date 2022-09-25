package main

import (
	"testing"
)

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		p    int
		want int
	}{
		{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4, 5},
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9, 4},
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9, 0},
		{[]int{1, 7, 2, 3, 5, 4, 8, 6}, 1, 6},
	} {
		got := tc.p
		if got != tc.want {
			t.Errorf("ERR: int(%b): got = %v, want = %v", tc.p, got, tc.want)
		}
	}
}
