package main

import "fmt"

func sumd(n, sum int) {
	sum += n % 10
	n /= 10
	if n == 0 {
		fmt.Println(sum)
		return
	}
	sumd(n, sum)

}

func main() {
	var n, sum int
	fmt.Scan(&n)
	sumd(n, sum)
}
