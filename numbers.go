package main

import "fmt"

func allNum(i, n int) {
	fmt.Println(i)
	if i < n {
		allNum(i+1, n)
	}
	if i > 0 {
		fmt.Println(i)
		i--
	}
}

// func allNumOld(n int) {
// 	if n > 0 {
// 		fmt.Println(n)
// 		allNumOld(n - 1)
// 	}
// }
