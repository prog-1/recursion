package main

import "fmt"

func isPower(n int) {
	if n == 1 {
		fmt.Println("Is power of 2")
		return
	}
	if n%2 != 0 {
		fmt.Println("Is not power of 2")
		return
	} else {
		n = n / 2
		isPower(n)
	}

}
func main() {
	var n int
	fmt.Scan(&n)
	isPower(n)
}
