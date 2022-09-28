package main

import (
	"testing"
)

type Input struct {
	s []int
	p int
}

func areLess(s []int, p int, t *testing.T) {
	for _, el := range s {
		if el > p {
			t.Errorf("%v > %v", el, p)
		}
	}
}

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input Input
		want  int
	}{
		{"no repeated elements", Input{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4}, 5},
		{"repeated elements", Input{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9}, 4},
		{"empty left side", Input{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			
			input := tc.input

			got := rearrange(input.s, input.p)

			areLess(input.s[:got], input.s[got], t)

			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
