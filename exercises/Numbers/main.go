package main

import "fmt"

func num1(n int, s []int) []int {
	if n < len(s)+1 {
		return s
	}
	s = append(s, len(s)+1)
	return num1(n, s)
}
func num2(n int, s []int) []int {
	if n == 0 {
		return s
	}
	s = append(s, n)
	n--
	return num2(n, s)
}
func num(n int) (s1, s2 []int) {
	var s []int
	if n < 1 {
		return nil, nil
	}
	return num1(n, s), num2(n, s)
}
func main() {
	fmt.Println(num1(8, []int{}))
}
