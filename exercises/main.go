package main

import (
	"fmt"
	"math"
)

func partition(s []int, pivot int) (i int) {
	l, r, p := 0, len(s)-1, s[pivot]
	for {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l >= r {
			return l
		}
		s[l], s[r] = s[r], s[l]

	}
}

func fromNtoOne(n int) {
	if n <= 0 {
		return
	}
	fromNtoOne(n - 1)
	fmt.Println(n)
	fromNtoOne(n - 1)
	//Time / Space complexity O(n) - linear func
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
	//Time / Space complexity O(n) - linear func
}
func isPowerOfTwo(n int) {
	n = n / 2
	if n == 2 {
		fmt.Println("TRUE")
	} else if n > 2 {
		isPowerOfTwo(n)
	} else {
		fmt.Println("FALSE")
	}
	//Time / Space complexity O(n) - logarithmic func, beacuse we divide by two every iteration i suppose
}
func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumOfDigits(n/10)
	//Time / Space complexity O(n) - logarithmic func, same as previous but i am not sure beacuse of %

}
func sliceMax(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	return int(math.Max(float64(s[0]), float64(sliceMax(s[1:]))))
	//Time / Space complexity ? can even we count complexity of something thats is not really  an algorithm

}
