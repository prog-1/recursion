package main

import (
	"reflect"
	"testing"
)

func ExampleNumbers1() {
	numbers(10)

	//Output:
	//1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10

}

func ExampleNumbers2() {
	numbers(4)

	//Output:
	//1
	//2
	//3
	//4
}

var fibonacciNumbers = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040}

func TestFibanachi(t *testing.T) {
	for i := range fibonacciNumbers {
		if got := fibonacci(i + 1); !reflect.DeepEqual(got, fibonacciNumbers[:i+1]) {
			t.Errorf("TestFibanachi(%v) got %v, want %v", i, got, fibonacciNumbers[:i])
		}
	}
}

func TestIsPowerOf2(t *testing.T) {
	for _, tc := range []struct {
		input float64
		want  bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{5, false},
		{6, false},
		{8, true},
		{12, false},
	} {
		if got := isPowerOf2(tc.input); got != tc.want {
			t.Errorf("isPowerOf2(%v)   got = %v, want = %v", tc.input, got, tc.want)
		}
	}
}

func TestNumSum(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  int
	}{
		{123, 6},
		{321, 6},
		{3201, 6},
		{3210, 6},
		{67, 13},
		{91, 10},
		{12, 3},
	} {
		if got := numSum(tc.input); got != tc.want {
			t.Errorf("numSum(%v)   got = %v, want = %v", tc.input, got, tc.want)
		}
	}
}

func TestMaxNum(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		wantMax1 int
		wantMax2 int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{2, 1, 0}, 2, 1},
		{[]int{2, 0, 5}, 5, 2},
	} {
		if max1, max2 := maxNum(tc.input); max1 != tc.wantMax1 || max2 != tc.wantMax2 {
			t.Errorf("numSum(%v) got = %v, %v, want = %v, %v", tc.input, max1, max2, tc.wantMax1, tc.wantMax2)
		}
	}
}

func shallowEqual(a []int, b []int) bool {
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
		if m[v] == -1 {
			return false
		}
	}
	return true
}

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		slice     []int
		pivot     int
		wantSlice []int
		wantIndex int
	}{
		{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4, []int{2, 4, 1, 0, 3, 5, 9, 7, 8, 6}, 5},
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9, []int{2, 5, 1, 5, 6, 9, 6, 7, 6, 9}, 4},
		{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9, []int{0, 6, 1, 7, 5, 9, 6, 5, 2, 9}, 0},
	} {
		gotSlice := make([]int, len(tc.slice))
		copy(gotSlice, tc.slice)
		if gotIndex := rearrange(gotSlice, tc.pivot); !shallowEqual(gotSlice[:gotIndex], tc.wantSlice[:tc.wantIndex]) {
			t.Errorf("rearange(%v, %v) = %v, %v, want %v, %v", tc.slice, tc.pivot, gotSlice, gotIndex, tc.wantSlice, tc.wantIndex)
		}
	}
}
