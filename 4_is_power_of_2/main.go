package main

import "fmt"

func IsPowerOf2(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 != 0 || n == 0 {
		return false
	}
	return IsPowerOf2(n / 2)
}

func main() {
	fmt.Println(IsPowerOf2(32))
	fmt.Println(IsPowerOf2(31))
	fmt.Println(IsPowerOf2(48))
}

// Time complexity: O(log(n))
// Space complexity: O(log(n))
