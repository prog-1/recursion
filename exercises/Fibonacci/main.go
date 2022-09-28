package main

import "fmt"

func fibonacci(n, a, b int, s []int) []int {
	if n > 0 {
		s = append(s, a)
		return fibonacci(n-1, b, a+b, s)
	}
	return s
}

func main() {
	var n int
	var s []int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n, 0, 1, s))
}
