package main

import (
	"testing"
)

func TestPower(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	} {
		if got := Power(tc.input); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
