package main

import "fmt"

func numbers(n, a int) {
	if n == 0 {
		return
	}
	fmt.Println(a)
	numbers(n-1, a+1)
}
func numbers2(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	numbers2(n - 1)
}

func main() {
	var n, a int
	fmt.Scan(&n)
	numbers(n, a)
	numbers2(n)
}
