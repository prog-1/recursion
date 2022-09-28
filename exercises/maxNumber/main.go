package main

import (
	"fmt"
)

func max(s []int, i, max1, max2 int) {
	if i == len(s) {
		fmt.Println(max1, max2)
		return
	}
	if max2 < s[i] && s[i] <= max1 {
		max2 = s[i]
	} else if max1 < s[i] {
		max1, max2 = s[i], max1
	}
	max(s, i+1, max1, max2)
}

// Time complexity: O(n)
// Space complexity: O(n)
