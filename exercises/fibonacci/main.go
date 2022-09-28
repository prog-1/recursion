package main

import "fmt"

func fib(Fn, F0, F1, i int) {
	if i == Fn {
		return
	}
	fmt.Print(F0+F1, " ")
	F0, F1 = F1, F0+F1
	i++
	fib(Fn, F0, F1, i)
}

func main() {
	var Fn int
	F0, F1, i := 0, 1, 0
	fmt.Print("Enter your Fn:  ")
	fmt.Scan(&Fn)
	fmt.Print("Fibonacci sequency - ")
	fmt.Print(F0, " ")
	fib(Fn, F0, F1, i)
}

// Time complexity: O(n)
// Space complexity: O(n)
