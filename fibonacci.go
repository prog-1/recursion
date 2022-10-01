package main

import "fmt"

func main23() {
	fibonacci(0)
}

//my interpritation of fibonacci using recursion
func fibonacci(n int) {
	var n1, n2 = 1, 0
	fibo(n1, n2, n)
}

func fibo(n1, n2, n int) {
	if n2 <= n {
		fmt.Println(n2)
		fibo(n1+n2, n1, n)
	}
}

//Time complexity: O(n) - calling recursion till we will not reach value of N
//Space complexity: O(1) - creating 2 variables at the start and then creating 3 input variables for each recursive function
