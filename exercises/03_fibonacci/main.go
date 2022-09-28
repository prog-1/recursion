package main

import "fmt"

func fibonacci(n, el1, el2 int) {
	if n != 0 {
		fmt.Printf("%v\n", el1)
		sum_el := el1 + el2
		el1 = el2
		el2 = sum_el
		n--
		fibonacci(n, el1, el2)
	}
}

func main() {
	var n, el1, el2 int
	el1 = 0
	el2 = 1
	fmt.Scan(&n)
	fibonacci(n, el1, el2)
}
