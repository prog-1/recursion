package main

import "fmt"

func maxnumber(s []int, max, a int) int {
	if len(s) == 1 {
		return max
	}
	if a < len(s) {
		if max < s[a] {
			max = s[a]
		}
		return maxnumber(s, max, a+1)
	}
	return max
}

func maxnumber2nd(s []int, max, max2, a int) int {
	if len(s) == 1 {
		return max
	}
	if max < max2 {
		max, max2 = max2, max
	}
	if a < len(s) {
		if max < s[a] {
			max2 = max
			max = s[a]
		} else if max2 < s[a] {
			max2 = s[a]
		}
		return maxnumber2nd(s, max, max2, a+1)
	}
	return max2
}

func main() {
	s := []int{1000, 1001, -100, 1}
	fmt.Println(maxnumber(s, s[0], 1))
	fmt.Println(maxnumber2nd(s, s[0], s[1], 2))
}
