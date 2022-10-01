package main

import "testing"

func TestPartition(t *testing.T) {
	for _, tc := range []tests{
		{"#1: no repeated elements", Input{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4}, 5},
		{"#2: repeated elements", Input{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9}, 4},
		{"#3: empty left side", Input{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {

			input := tc.input

			got := partition(input.s, input.p)

			check(input.s[:got], input.s[got], t)

			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func check(s []int, p int, t *testing.T) {
	//checking if each el in left side is smaller then pivot
	for _, elem := range s {
		if elem > p {
			t.Errorf("%v > %v", elem, p)
		}
	}
}

type tests struct {
	name  string
	input Input
	want  int
}

type Input struct {
	s []int
	p int
}
