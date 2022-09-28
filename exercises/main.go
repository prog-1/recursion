package main

import "fmt"

func numbers1(n int) {

	fmt.Println(n)

	if n > 1 {
		numbers1(n-1)
	}

}

func numbers2(n int) {

	if n > 1 {
		numbers2(n-1)
	}

	fmt.Println(n)

}


func main() {
	numbers1(10)
}
