package main

import (
	"fmt"
	"math"
)

func rearrange(s []int, pivot int) (i int) {
	l, r := 0, len(s)-1
	p := s[pivot]
	for {
		for s[l] < p {
			l++
		}
		for s[r] >= p {
			r--
		}
		if l >= r {
			return r + 1
		}
		if l == pivot {
			s[l], s[l+1] = s[l+1], s[l]
			pivot++
		}
		s[l], s[r] = s[r], s[l]

		l, r = l+1, r-1
	}
}

// Time complexity: O(n)
// Space complexity: O(n)
func numbers(n int) {
	if n < 1 {
		return
	}
	numbers(n - 1)
	fmt.Println(n)
}

// Time complexity: O(n)
// Space complexity: O(n)
func fibonacci(n int) []int {
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 1}
	}
	a := fibonacci(n - 1)
	return append(a, a[len(a)-1]+a[len(a)-2])
}

// Time complexity: O(log n)
// Space complexity: O(log n)
func isPowerOf2(n float64) bool {
	if n == 2 || n == 1 {
		return true
	} else if n < 2 {
		return false
	}
	return isPowerOf2(n / 2)
}

// Time complexity: O(len(n))
// Space complexity: O(len(n))
func numSum(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + numSum(n/10)
}

// Time complexity: O(len(n))
// Space complexity: O(len(n))
func maxNum(s []int) (int, int) {
	if len(s) == 1 {
		return s[0], math.MinInt
	}
	max1, max2 := maxNum(s[1:])
	if max1 < s[0] {
		return s[0], max1
	} else if max2 < s[0] {
		return max1, s[0]
	}
	return max1, max2

}

func main() {
	fmt.Println(fibonacci(30))
}
