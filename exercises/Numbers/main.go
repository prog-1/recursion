package main

import "fmt"

func num1(n, i int) {
	if n < i {
		return
	}
	fmt.Println(i)
	i++
	num1(n, i)
}
func num2(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	num2(n)
}
func main() {
	var n int
	i := 1
	fmt.Scan(&n)
	num1(n, i)
	num2(n)
}
