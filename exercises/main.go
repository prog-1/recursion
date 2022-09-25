package main

import "fmt"

func partition(s []int, pivot int) (i int) {

	p := s[pivot]
	for {
		fmt.Println(s)
		l, r := 0, len(s)-1

		for s[l] <= p {
			l++
		}
		for s[r] > p && r > 0 {
			r--
		}

		if l >= r && r > 0 {
			fmt.Println(s)
			return r
		} else if l >= r && r <= 0 {
			return r
		}
		s[l], s[r] = s[r], s[l]

	}

}

func main() {
	s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}
	i := partition(s, 9)
	fmt.Println("Output: ", i, s)
	fmt.Println("Left side: ", s[:i])
	fmt.Println("Right side: ", s[i:])
}
