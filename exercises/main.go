package main

import "fmt"

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

func main() {
	s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}
	fmt.Println(rearrange(s, 9), s) // Output: 4 [2 5 1 5 6 9 6 7 6 9]
}
