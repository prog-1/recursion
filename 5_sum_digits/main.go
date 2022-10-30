package main

import "fmt"

func SumDigits(n int) int {
	if n == 0 {
		return 0
	}
	return SumDigits(n/10) + n%10
}

func main() {
	fmt.Println(SumDigits(561))
}

// Time complexity: O(log(n))
// Space complexity: O(log(n))
