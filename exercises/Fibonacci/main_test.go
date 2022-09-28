package main

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
		a    int
		b    int
		s    []int
		want []int
	}{
		{"1", 1, 0, 1, []int{}, []int{0, 1}},
		{"2", 8, 0, 1, []int{}, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{"3", -1, 0, 1, []int{}, []int{}},
		{"4", 5, 0, 1, []int{}, []int{0, 1, 1, 2, 3}},
		{"5", 0, 0, 1, []int{}, []int{0}},
		{"6", 2, 0, 1, []int{}, []int{0, 1, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := fibonacci(tc.n, tc.a, tc.b, tc.s)
			if reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
