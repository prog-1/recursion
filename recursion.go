package main

import "fmt"

func partition(s []int) int {
	l, r := 0, len(s)-1
	p := s[r/2]
	for {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l >= r {
			return r
		}
		s[l], s[r] = s[r], s[l]
		l, r = l+1, r-1
	}
}

// For each call:
// - partition: space complexity O(1)
// - for the quick sort itself (one call):
//     - best case:  O(1) * O(LogN) = O(LogN)
//     - worst case: O(1) * O(N)    = O(N)
//
// Stack, LIFO (Last In First Out)
// In:  1 2 3 4 5
// Out: 5 4 3 2 1
// Stack contains: Local variables, Arguments, Return address
func QuickSort(s []int) {
	if len(s) < 2 {
		return
	}
	p := partition(s) + 1 // n ->  l r   1/n-1
	// Recursion.
	QuickSort(s[:p])
	QuickSort(s[p:])
}

// 1 2   3 4   5 6   7 8
// Ideal case depth: LogN
// Time complexity: O(NLogN)
//
// Worst case depth: N
// Time complexity: O(N^2)
//
// Pivot is everything!!!
//
// 25% 75% NLogN
func main() {
	s := []int{4, 3, 17, 33, 101, 17, 6, 2}
	QuickSort(s)
	fmt.Println(s)
}
