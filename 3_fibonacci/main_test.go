package main

import (
	"fmt"
	"reflect"
	"testing"
)

var numbers = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}

func TestFibonacci(t *testing.T) {
	for _, tc := range []struct {
		input int
	}{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{12},
		{13},
		{18},
		{20},
		{25},
		{30},
		{56},
		{99},
		{100},
		{143},
		{144},
		{145},
		{228},
		{4200},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			want := numbers
			for i := range numbers {
				if numbers[i] > tc.input {
					want = numbers[:i]
					break
				}
			}
			if got := Fibonacci(tc.input); !reflect.DeepEqual(got, want) {
				t.Errorf("Fibonacci(%v) = %v, want %v", tc.input, got, want)
			}
		})
	}
}
