package main

import "fmt"

func rearrange(s []int, pivot int) (i int) {
	l, r := 0, len(s)-1
	p := s[pivot]
	for {
		for ; s[l] < p; l++ {
		}
		for ; s[r] >= p && r > 0; r-- {
		}
		if l >= r {
			if r == 0 {
				return r
			} else {
				return r + 1
			}
		}

		s[l], s[r] = s[r], s[l]

		l++
		r--
	}
}

func main() {
	p := 9
	fmt.Println("Pivot:", p)
	s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}
	fmt.Println("Initial:", s)
	fmt.Println()
	i := rearrange(s, p)
	fmt.Println("Index:", i)
	fmt.Println("Processed:", s)
	fmt.Println("s[:i]:", s[:i])
	fmt.Println("s[i:]:", s[i:])
}
