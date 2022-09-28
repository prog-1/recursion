package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	for _, tc := range []struct {
		name         string
		n            int
		want1, want2 []int
	}{
		{"1", 1, []int{1}, []int{1}},
		{"2", 2, []int{1, 2}, []int{2, 1}},
		{"3", 0, nil, nil},
		{"4", -1, nil, nil},
		{"5", 5, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"6", 8, []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got1, got2 := num(tc.n)
			if !reflect.DeepEqual(got1, tc.want1) || !reflect.DeepEqual(got2, tc.want2) {
				t.Errorf("got1 = %v, want1 = %v ; got2 = %v, want2 = %v", got1, tc.want1, got2, tc.want2)
			}
		})
	}
}
