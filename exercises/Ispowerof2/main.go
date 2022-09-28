package main

import "fmt"

func isPower(n int) bool {
	if n == 2 {
		fmt.Println("Is power of 2")
		return true
	}
	if n%2 != 0 {
		fmt.Println("Is not power of 2")
		return false
	} else {
		n = n / 2
		return isPower(n)
	}

}
func main() {

}
