package main

import "fmt"

func number(i int) {
	if i <= 0 {
		return
	}
	fmt.Printf("%v\n", i)
	if i >= 1 {
		number(i - 1)
		fmt.Printf("%v\n", i)
	}
}

//time complexity: O(n)
//space complexity: O(n)

func main() {
	var i int
	fmt.Scan(&i)
	number(i)
}
