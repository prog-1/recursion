package main

import "fmt"

func main() {
	s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}
	fmt.Println("Initial s: ", s)
	i := partition(s, 9)
	fmt.Println("I: ", i)
	fmt.Println("Recieved s: ", s)
	fmt.Println("Left side: ", s[:i])  // < i
	fmt.Println("Right side: ", s[i:]) // other
}

func partition(s []int, pivot int) (i int) {
	l, r := 0, len(s)-1
	p := s[pivot]
	for {
		for s[l] < p {
			l++
		}
		for s[r] >= p && r > 0 { // 3rd works with validation of r on 0
			r--
		}
		if l >= r && r > 0 {
			return r + 1
		} else if l >= r && r <= 0 {
			return r
		}
		s[l], s[r] = s[r], s[l]

	}
}
