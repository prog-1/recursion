package main

import "fmt"

func ispowerof2(n int) bool {
	if n%2 != 0 {
		return false
	}
	if n == 2 || n == 1 {
		return true
	}
	return ispowerof2(n / 2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(ispowerof2(n))
}
