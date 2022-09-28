package main

import "fmt"

func sumdigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumdigits(n/10)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumdigits(n))
}
