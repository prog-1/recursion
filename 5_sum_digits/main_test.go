package main

import (
	"fmt"
	"testing"
)

func TestSumDigits(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{5, 5},
		{9, 9},
		{10, 1},
		{11, 2},
		{12, 3},
		{18, 9},
		{19, 10},
		{20, 2},
		{24, 6},
		{55, 10},
		{89, 17},
		{100, 1},
		{123, 6},
		{561, 12},
		{345678, 33},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := SumDigits(tc.input); got != tc.want {
				t.Errorf("SumDigits(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
