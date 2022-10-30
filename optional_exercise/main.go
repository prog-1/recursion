package main

import "fmt"

func partition(s []int) (i int) {
	p := s[(len(s) - 1)]
	for j := 0; j < len(s); j++ {
		if s[j] < p {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[(len(s)-1)] = s[(len(s)-1)], s[i]
	return i
}

// For each call:
// - partition: space complexity O(1)
// - for the quick sort itself (one call):
//     - best case:  O(1) * O(LogN) = O(LogN)
//     - worst case: O(1) * O(N)    = O(N)
//
// Stack, LIFO (Last In First Out)
// In:  1 2 3 4 5
// Out: 5 4 3 2 1
// Stack contains: Local variables, Arguments, Return address

func QuickSort(s []int) (result []int) {
	s1, s2 := make([]int, len(s)), make([]int, len(s))
	copy(s1, s)
	copy(s2, s)
	for len(s1) >= 2 || len(s2) >= 2 {
		if len(s1) >= 2 {
			p := partition(s1)
			s1 = s1[:p]
			fmt.Println("s1:", s1)
		}
		result = append(result, s1...)
		if len(s2) >= 2 {
			p := partition(s2)
			s2 = s2[p+1:]
			fmt.Println("s2:", s2)
		}
		result = append(result, s2...)
	}
	return
}

// 1 2   3 4   5 6   7 8
// Ideal case depth: LogN
// Time complexity: O(NLogN)
//
// Worst case depth: N
// Time complexity: O(N^2)
//
// Pivot is everything!!!
//
// 25% 75% NLogN

func main() {
	s := []int{4, 3, 17, 33, 101, 17, 6, 2}
	fmt.Println(QuickSort(s))
	//fmt.Println(s)
}
