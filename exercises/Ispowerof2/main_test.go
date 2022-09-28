package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
		want bool
	}{
		{"1", 8, true},
		{"2", 1, false},
		{"3", -1, false},
		{"4", 5, false},
		{"5", 1024, true},
		{"6", 128, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := isPower(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
