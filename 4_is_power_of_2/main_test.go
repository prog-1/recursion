package main

import (
	"fmt"
	"testing"
)

func TestIsPowerOf2(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  bool
	}{
		{0, false},
		{1, true},
		{-1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, false},
		{8, true},
		{9, false},
		{10, false},
		{15, false},
		{16, true},
		{20, false},
		{32, true},
		{-32, false},
		{64, true},
		{-64, false},
		{65, false},
		{-65, false},
		{66, false},
		{-66, false},
		{128, true},
		{256, true},
		{512, true},
		{1024, true},
		{1524, false},
		{2048, true},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := IsPowerOf2(tc.input); got != tc.want {
				t.Errorf("IsPowerOf2(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
