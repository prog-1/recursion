package main

import (
	"testing"
)

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		name  string
		s     []int
		pivot int
		want  int
	}{
		{"no repeated elements", []int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4, 5},
		{"repeated elements", []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9, 4},
		{"empty left side", []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := rearrange(tc.s, tc.pivot)
			areLess(tc.s[:got], tc.s[got], t)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func areLess(s []int, p int, t *testing.T) {
	for _, e := range s {
		if e > p {
			t.Errorf("%v > %v", e, p)
		}
	}
}
