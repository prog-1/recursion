package main

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	t.Run("Fibonacci test", func(t *testing.T) {
		want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
		got := fibonacci(10)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input int
		want  bool
	}{
		{"Is power of two", 8, true},
		{"Not power of two", 9, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := isPowerOfTwo(tc.input)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	t.Run("Digit sum test", func(t *testing.T) {
		want := 4
		got := sumDigits(121)
		if want != got {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}

func TestMaxNum(t *testing.T) {
	t.Run("Max num test", func(t *testing.T) {
		want := []int{3, 2}
		input := []int{1, 2, 3}
		got := maxNum(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}
