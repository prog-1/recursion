package main

import "fmt"

func power(n int) {
	if n == 1 {
		fmt.Println("Is power of two")
		return
	} else if n <= 0 || n%2 != 0 {
		fmt.Println("Is not power of 2")
		return
	} else if n >= 2 {
		n /= 2
		power(n)
	}
}

//time complexity: O(logn)
//space complexity: O(logn)

func main() {
	var n int
	fmt.Scan(&n)
	power(n)
}
