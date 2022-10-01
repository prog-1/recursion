package main

import "fmt"

func main44() {
	fmt.Println(sumDigits(1))
}

func sumDigits(n int) int {
	var sum, remainder int
	return sDigits(n, sum, remainder)
}

func sDigits(n, sum, remainder int) int {
	if n > 0 {
		remainder = n % 10
		n /= 10
		sum += remainder
		return sDigits(n, sum, remainder)
	} else {
		return sum
	}
}

//Time complexity: O(n)
//Space complexity: O(n)
