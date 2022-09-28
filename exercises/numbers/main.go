package main

import "fmt"

func print(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	print(n - 1)
	fmt.Println(n)
}

// Time complexity: O(n)
// Space complexity: O(n)
