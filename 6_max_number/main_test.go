package main

import (
	"fmt"
	"testing"
)

func TestMaxNumber(t *testing.T) {
	for _, tc := range []struct {
		input         []int
		wantmax       int
		wantsecondmax int
	}{
		{[]int{}, -1, -1},
		{[]int{0}, 0, -1},
		{[]int{1}, 1, -1},
		{[]int{1, 1}, 1, 1},
		{[]int{0, 1}, 1, 0},
		{[]int{1, 0}, 1, 0},
		{[]int{2, 1}, 2, 1},
		{[]int{1, 2}, 2, 1},
		{[]int{2, 2}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{4, 2, 3, 1}, 4, 3},
		{[]int{1, 2, 6, 9, 3, 1, 10, 15, 7, 11}, 15, 11},
		{[]int{1, 2, 6, 9, 3, 1, 10, 15, 7, 15}, 15, 15},
		{[]int{1, 2, 6, 9, 3, 1, 10, 15, 7, 16}, 16, 15},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2, 1},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if gotmax, gotsecondmax := MaxNumber(tc.input); gotmax != tc.wantmax || gotsecondmax != tc.wantsecondmax {
				t.Errorf("MaxNumber(%v) = %v %v, want %v %v", tc.input, gotmax, gotsecondmax, tc.wantmax, tc.wantsecondmax)
			}
		})
	}
}
