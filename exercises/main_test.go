package main

import (
	"testing"
)

type Input struct {
	s []int
	p int
}

func areLess(s []int, p int, t *testing.T) {
	for _, el := range s {
		if el > p {
			t.Errorf("%v > %v", el, p)
		}
	}
}

func TestRearrange(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input Input
		want  int
	}{
		{"case1", Input{[]int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}, 4}, 5},
		{"case2", Input{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}, 9}, 4},
		{"case3", Input{[]int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}, 9}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {

			input := tc.input

			got := partitions(input.s, input.p)

			areLess(input.s[:got], input.s[got], t)

			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

// func Testfibonacci(t* testing.T){
// 	for _, tc := range []struct{
// 		n int
// 		want int
// 	}{
// 		{0, 1, 1, 2, 3, 5, 8}
// 	}
// 	}
// }
func TestIspowerof2(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{512, true},
		{128, true},
		{9, false},
		{3, false},
		{64, true},
		{31, false},
	} {
		if got := powerof2(tc.n); got != tc.want {
			t.Errorf("powerof2(%v)   got = %v, want = %v", tc.n, got, tc.want)
		}
	}
}

func TestSumdigits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		sum  int
		want int
	}{
		{333, 0, 9},
		{321, 0, 6},
		{21, 0, 3},
		{10000, 0, 1},
		{420, 0, 6},
		{228, 0, 12},
	} {
		if got := sumdigits(tc.n, tc.sum); got != tc.want {
			t.Errorf("numSum(%v)   got = %v, want = %v", tc.n, got, tc.want)
		}
	}
}

func TestMaxnum(t *testing.T) {
	for _, tc := range []struct {
		n     []int
		i     int
		max   int
		max2  int
		wmax  int
		wmax2 int
	}{
		{[]int{1000, 1001, -100, 1}, 0, 0, 0, 1001, 1000},
		{[]int{1, -3, 99, 675}, 0, 0, 0, 675, 99},
		{[]int{10, 11, -100, 1}, 0, 0, 0, 11, 10},
	} {
		if max, max2 := maxnum(tc.n, tc.i, tc.max, tc.max2); max != tc.wmax || max2 != tc.wmax2 {
			t.Errorf("numSum(%v) got = %v, %v, want = %v, %v", tc.n, max, max2, tc.wmax, tc.wmax2)
		}
	}
}
