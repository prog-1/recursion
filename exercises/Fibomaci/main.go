package main

import (
	"fmt"
)

var s = []int{0, 1}
var cnt = 1

func cycle(n int) []int {
	s = append(s, s[len(s)-1]+s[len(s)-2])
	cnt++
	if cnt != n {
		cycle(n)
	}
	return s
}
func fib(n int) []int {
	if n < 1 {
		return nil
	} else if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}
	return cycle(n)
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n))
}
