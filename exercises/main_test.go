package main

import (
	"reflect"
	"testing"
)

type want struct {
	i int
	s []int
}

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		name string
		str  []int
		p    int
		want want
	}{
		{"Example 1", []int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4, want{5, []int{2, 4, 1, 0, 3, 5, 9, 7, 8, 6}}},
		{"Example 2", []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9, want{4, []int{2, 5, 1, 5, 6, 9, 6, 7, 6, 9}}},
		{"Example 3", []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9, want{0, []int{0, 6, 1, 7, 5, 9, 6, 5, 2, 9}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := partition(tc.str, tc.p)
			if got != tc.want.i || !reflect.DeepEqual(tc.str, tc.want.s) {
				t.Errorf("got = %v,%d want = %v", got, tc.str, tc.want)
			}
		})
	}
}
