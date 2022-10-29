package main

import "fmt"

func Fibonacci(n int) []int {
	return fib(n, 0, 1, nil)
}

func fib(n, a, b int, result []int) []int {
	result = append(result, a)
	if b > n {
		return result
	}
	return fib(n, b, a+b, result)
}

func main() {
	fmt.Println(Fibonacci(13))
	fmt.Println(Fibonacci(15))
}

// Time complexity: O(n)
// Space complexity: O(n)
