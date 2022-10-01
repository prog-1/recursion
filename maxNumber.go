package main

import "fmt"

func main2() {
	s := []int{ /*7, 1, 2, 3, 5, 4*/ 0, 1}
	maxNumber(s)
}

// Time Complexity: O(N^2) - Worst case
// Space Complexity: O(logN)


func maxNumber(s []int) {
	if len(s)-1 > 0 {
		quickSort(s) // programmers must be lazy ;)
		fmt.Println("First max number:", s[len(s)-1], "Second max number:", s[len(s)-2])
	} else {
		fmt.Println("First max number:", s[len(s)-1])
	}
}

func quickSort(s []int) {
	if len(s) < 2 {
		return
	}
	p := partition(s) + 1
	quickSort(s[:p])
	quickSort(s[p:])
}

func partition(s []int) (i int) {
	l, r := 0, len(s)-1
	p := s[r/2]
	for {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l >= r {
			return l
		}
		s[l], s[r] = s[r], s[l]
		//l++
		r--
	}
}
