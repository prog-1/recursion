package main

import "fmt"

func max(s []int) (int, int) {
	var min, max1, max2 int

	if len(s) > 0 {
		min = s[0]
		max1, max2 = max(s[1:])
	}
	if min > max1 {
		return min, max1
	} else if min > max2 {
		return max1, min
	} else {
		return max1, max2
	}
}

//time complexity: O(n)
//space complexity: O(n)

func main() {
	s := []int{23, 10, 15, 0, 9, 20, 10}
	fmt.Println(max(s))
}
