package main

import "fmt"

func num1(n, i int) {
	if n < i {
		return
	}
	fmt.Print(i, " ")
	i++
	num1(n, i)
}
func num2(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	n--
	num2(n)
}
func main() {
	var n int
	i := 1
	fmt.Print("Enter your n:  ")
	fmt.Scan(&n)
	fmt.Print("Ascending order - ")
	num1(n, i)
	fmt.Print("Descending order - ")
	num2(n)
}

// Time complexity: O(n)
// Space complexity: O(n)
