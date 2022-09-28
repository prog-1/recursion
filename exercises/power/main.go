package main

import "fmt"

func Power(n int) {
	if n == 1 {
		fmt.Println("Is power of 2.")
		return
	}
	if n%2 != 0 {
		fmt.Println("Is not power of 2.")
		return
	} else {
		n = n / 2
		Power(n)
	}
}
func main() {
	var n int
	fmt.Print("Enter your n:  ")
	fmt.Scan(&n)
	Power(n)
}
