package main

import "fmt"

func isPowerOf2(n int) {
	if n%2 == 1 && n != 1 {
		fmt.Println("NO")
		return
	} else if n == 1 {
		fmt.Println("YES")
		return
	}
	isPowerOf2(n / 2)
}

// Time complexity: O(log(n))
// Space complexity: O(log(n))
