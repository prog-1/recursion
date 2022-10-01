package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOf2(1))
}
func isPowerOf2(n int) bool {
	if n < 2 {
		return false
	}
	logN := math.Log2(float64(n))   //taking logarithm with base of 2 from N
	return logN == math.Round(logN) //if number is integer (has no leftover)

}

//Time complexity: O(1) - no loops, no recursion, one calculation
//Space complexity: O(1) - only input variable and logN variable
