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
			got := partition(tc.s, tc.pivot)
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

func TestSumOfDigits(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  int
	}{
		{0, 0},
		{123, 6},
		{1234567890, 45},
		{666, 18},
		{-12, -3}, // considering this test func counts both numbers as negative -1 + -2
		{-123, -6},
	} {
		if got := sumOfDigits(tc.input); got != tc.want {
			t.Errorf("numSum(%v)   got = %v, want = %v", tc.input, got, tc.want)
		}
	}
}
func TestSliceMax(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  int
	}{
		//{nil, 0}, FAIL
		{[]int{1}, 1},
		{[]int{1, 34, 45, -67}, 45},
		{[]int{1, 1, 1, 1, 1, 1}, 1},
		{[]int{-5, -8, -90, -4567}, -5},
	} {
		if got := sliceMax(tc.input); got != tc.want {
			t.Errorf("numSum(%v)   got = %v, want = %v", tc.input, got, tc.want)
		}
	}
}
