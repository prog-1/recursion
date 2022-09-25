package main

import (
	"testing"
)

func shallowEqual(a []int, b []int) bool {
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
		if m[v] == -1 {
			return false
		}
	}
	return true
}

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		slice     []int
		pivot     int
		wantSlice []int
		wantIndex int
	}{
		{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4, []int{2, 4, 1, 0, 3, 5, 9, 7, 8, 6}, 5},
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9, []int{2, 5, 1, 5, 6, 9, 6, 7, 6, 9}, 4},
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9, []int{0, 6, 1, 7, 5, 9, 6, 5, 2, 9}, 0},
	} {
		gotSlice := make([]int, len(tc.slice))
		copy(gotSlice, tc.slice)
		if gotIndex := rearrange(gotSlice, tc.pivot); !shallowEqual(gotSlice[:gotIndex], tc.wantSlice[:tc.wantIndex]) {
			t.Errorf("rearange(%v, %v) = %v, %v, want %v, %v", tc.slice, tc.pivot, gotSlice, gotIndex, tc.wantSlice, tc.wantIndex)
		}
	}
}
