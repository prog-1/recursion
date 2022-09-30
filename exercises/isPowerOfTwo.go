package main

import (
	"math"
)

func isWhole(n float64) bool {
	return n == math.Round(n)
}

func isPowerOfTwo(n int) bool {
	log := math.Log2(float64(n))
	return isWhole(log)
	// Time complexity O(1)
	// Space complexity O(1)
}

// func main() {
// 	n := 16

// 	if isPowerOfTwo(n) {
// 		fmt.Println("YEAH")
// 	} else {
// 		fmt.Println("NOPE")
// 	}
// }
