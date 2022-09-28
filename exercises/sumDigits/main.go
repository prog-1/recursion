package main

import (
	"fmt"
)

func sumDigits(n, sum int) {
	if n == 0 {
		fmt.Println(sum)
		return
	}
	sum += n % 10
	sumDigits(n/10, sum)
}

func main() {
	sumDigits(123, 0)
	sumDigits(987, 0)
	sumDigits(12345, 0)
	sumDigits(12357, 0)
}

// Time complexity: O(n)
// Space complexity: O(n)
