package main

import (
	"fmt"
)

func rearrange(s []int, pivot int) (i int) {
	cnt, p := 0, s[pivot]
	j := 0
	for ; j < len(s); j++ {
		if s[j] == s[pivot] {
			cnt++
		}
		if s[j] < s[pivot] {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[len(s)-1] = s[len(s)-1], s[i]
	if cnt == 1 && p != 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == p {
				s[i], s[pivot+1] = s[pivot+1], s[i]
			}
		}
	}
	return i
}

func main() {
	s := []int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6, 5}
	fmt.Println(rearrange(s, 4), s) // Output: 5 [2 4 1 0 3 5 9 7 8 6]

	s4 := []int{4, 4, 4, 4}
	fmt.Println(rearrange(s4, 2), s4) // Output: 5 [2 4 1 0 3 5 9 7 8 6]

	s0 := []int{0}
	fmt.Println(rearrange(s0, 0), s0) // Output: 5 [2 4 1 0 3 5 9 7 8 6]

	s2 := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}
	fmt.Println(rearrange(s2, 9), s2) // Output: 4 [2 5 1 5 6 9 6 7 6 9]

	s3 := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}
	fmt.Println(rearrange(s3, 9), s3) // Output: 0 [0 6 1 7 5 9 6 5 2 9]
}
