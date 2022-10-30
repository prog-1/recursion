package main

import "fmt"

func MaxNumber(s []int) (int, int) {
	return Max(s, -1, -1)
}

func Max(s []int, max, secondmax int) (int, int) {
	if len(s) == 0 {
		return max, secondmax
	}
	if s[0] >= max {
		secondmax = max
		max = s[0]
	}
	if s[0] > secondmax && s[0] < max {
		secondmax = s[0]
	}
	return Max(s[1:], max, secondmax)
}

func main() {
	fmt.Println(MaxNumber([]int{1, 2, 6, 9, 3, 1, 10, 15, 7, 11}))
}

// Time complexity: O(n)
// Space complexity: O(n)
