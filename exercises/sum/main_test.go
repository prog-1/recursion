package main

import "fmt"

func TestSum(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  int
	}{
		{0, 0},
		{1234, 10},
		{4321, 10},
		{101, 2},
		{56, 11},
	} {
		if got := sumf(tc.input); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
