package main

import "fmt"

func Numbers(n int) ([]int, []int) {
	return Num(n, nil, nil)
}

func Num(n int, ascending, descenging []int) ([]int, []int) {
	if n == 0 {
		return ascending, descenging
	}
	ascending, descenging = Num(n-1, ascending, append(descenging, n)) // (2)
	return append(ascending, n), descenging                            // (1)
}

func main() {
	fmt.Println(Numbers(4))
}

// Time complexity: O(n)
// Space complexity: O(n)
