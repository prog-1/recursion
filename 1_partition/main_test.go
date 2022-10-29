package main

import (
	"fmt"
	"testing"
)

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		s     []int
		pivot int
		want  int
	}{
		{[]int{1, 2}, 1, 1},
		{[]int{3, 2, 4, 6}, 2, 2},
		{[]int{3, 2, 4, 6, 1}, 2, 3},
		{[]int{4, 4, 4, 4}, 2, 0},
		{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4, 5}, // no repeated elements
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9, 4}, // repeated elements
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9, 0}, // empty left side
	} {
		t.Run(fmt.Sprint(tc.s)+"_"+fmt.Sprint(tc.pivot), func(t *testing.T) {
			got := rearrange(tc.s, tc.pivot)
			CompareLeftAndRight(tc.s, got)
			if got != tc.want {
				t.Errorf("rearrange(%v, %v) = %v, want %v", tc.s, tc.pivot, got, tc.want)
			}
		})
	}
}

func CompareLeftAndRight(s []int, index int) error {
	for _, i := range s[:index] {
		for _, j := range s[index:] {
			if i >= j {
				panic("the returned slice is incorrect")
			}
		}
	}
	return nil
}
