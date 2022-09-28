package main

import "fmt"

func fib(n, x1, x2, cnt int) {
	if cnt == n {
		return
	}
	fmt.Print(" ", x1+x2)
	cnt++
	x1, x2 = x2, x1+x2
	fib(n, x1, x2, cnt)

}
func main() {
	var n, x1 int
	x2, cnt := 1, 1
	fmt.Scan(&n)
	fmt.Print(x1, " ", x2)
	fib(n, x1, x2, cnt)
}
