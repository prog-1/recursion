package main

import "fmt"

func sumf(n, sum int) {
	sum += n % 10
	n /= 10
	if n == 0 {
		fmt.Println(sum)
		return
	}
	sumf(n, sum)
}

func main() {
	var n, sum int
	fmt.Print("Enter your n:  ")
	fmt.Scan(&n)
	sumf(n, sum)
}

// Time complexity: O(len(n))
// Space complexity: O(len(n))
