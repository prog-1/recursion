package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	for _, tc := range []struct {
		name         string
		s            []int
		want1, want2 int
	}{
		{"1", []int{0, 1, 1, 2, 3, 5, 8, 13, 21}, 0, 21},
		{"2", []int{1}, 1, 1},
		{"3", []int{0, 1}, 0, 1},
		{"4", []int{0, 1, 32, -54, 21, 21}, -54, 32},
		{"5", []int{21, 32, 12, 32}, 12, 32},
		{"6", []int{-1234, 2123, 21, 23214546}, -1234, 23214546},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got1, got2 := minMax(tc.s)
			if got1 != tc.want1 || got2 != tc.want2 {
				t.Errorf("got1 = %v, want1 = %v ; got2 = %v, want2 = %v", got1, tc.want1, got2, tc.want2)
			}
		})
	}
}
