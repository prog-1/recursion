package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
		want int
	}{
		{"1", 12354, 15},
		{"2", 123456789, 45},
		{"3", -12345, 15},
		{"4", 0, 0},
		{"5", 1, 1},
		{"6", 128, 11},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := sumd(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
