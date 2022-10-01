package main

import "fmt"

func ExampleAllNum() {
	allNum(1, 10)

	//Output:
	//1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10
	//10
	//9
	//8
	//7
	//6
	//5
	//4
	//3
	//2
	//1
}

func ExampleFibonacci() {
	fibonacci(10)
	fibonacci(1)
	fibonacci(0)

	//Output:
	//0
	//1
	//1
	//2
	//3
	//5
	//8
	//0
	//1
	//1
	//0
}

func ExamplePowerOf2() {
	fmt.Println(isPowerOf2(256))
	fmt.Println(isPowerOf2(5))
	fmt.Println(isPowerOf2(65536))
	fmt.Println(isPowerOf2(0))
	fmt.Println(isPowerOf2(1))

	//Output:
	//true
	//false
	//true
	//false
	//false
}

func ExampleSumDigits() {
	fmt.Println(sumDigits(343))
	fmt.Println(sumDigits(11))
	fmt.Println(sumDigits(1))
	fmt.Println(sumDigits(0))
	fmt.Println(sumDigits(1000))

	//Output:
	//10
	//2
	//1
	//0
	//1
}

func ExampleMaxNumber() {
	maxNumber([]int{7, 1, 2, 3, 5, 4})
	maxNumber([]int{0})
	maxNumber([]int{0, 1})
	maxNumber([]int{-100, -50, -200})

	//Output:
	//First max number: 7 Second max number: 5
	//First max number: 0
	//First max number: 1 Second max number: 0
	//First max number: -50 Second max number: -100
}
