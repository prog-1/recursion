package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
		want []int
	}{
		{"1", 1, []int{0, 1}},
		{"2", 8, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{"3", -1, nil},
		{"4", 5, []int{0, 1, 1, 2, 3}},
		{"5", 0, []int{0}},
		{"6", 2, []int{0, 1, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := fib(tc.n)
			if reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
