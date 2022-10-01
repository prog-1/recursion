package main

import "fmt"

func sum(n, sum1 int) {
	if n > 0 {
		sum1 += n % 10
		n /= 10
		sum(n, sum1)
		if n == 0 {
			fmt.Println(sum1)
		}
	}
}

//time complexity: O(n)
//space complexity: O(n)

func main() {
	var n, sum1 int
	fmt.Scan(&n)
	sum(n, sum1)
}
