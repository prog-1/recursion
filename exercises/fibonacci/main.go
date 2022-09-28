package main

import (
	"fmt"
)

func fibonacci(n, n1, n2 int) {
	if n == 0 {
		return
	}
	fmt.Printf("%v ", n2)
	fibonacci(n-1, n2, n1+n2)
}

func main() {
	fibonacci(1, 0, 1)
	fmt.Println()
	fibonacci(2, 0, 1)
	fmt.Println()
	fibonacci(3, 0, 1)
	fmt.Println()
	fibonacci(4, 0, 1)
	fmt.Println()
	fibonacci(5, 0, 1)
	fmt.Println()
	fibonacci(6, 0, 1)
	fmt.Println()
	fibonacci(7, 0, 1)
}

// Time complexity: O(n)
// Space complexity: O(n)
