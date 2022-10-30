package main

import (
	"fmt"
)

func rearrange(s []int, pivot int) (i int) {
	p := s[pivot]
	for j := 0; j < len(s); j++ {
		if s[j] < p {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	return
}

func main() {
	s := []int{3, 2, 4, 6, 1}
	fmt.Println(rearrange(s, 2), s)
}
