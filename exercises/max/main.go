package main

import "fmt"

func Max(s []int, maxB, maxS int) {
	if len(s) == 0 {
		fmt.Println("Max - ", maxB, "Smaller max - ", maxS)
		return
	}
	if s[0] > maxB {
		maxB = s[0]
	}
	if len(s) == 2 {
		maxS = maxB
	}
	Max(s[1:], maxB, maxS)
}
func main() {
	maxB, maxS := 0, 0
	var count int
	fmt.Print("How many numbers do you want to enter? ")
	fmt.Scan(&count)
	var s []int
	for i := 0; i < count; i++ {
		var n int
		fmt.Printf("Enter number #%v: ", i+1)
		fmt.Scan(&n)
		s = append(s, n)
	}
	Max(s, maxB, maxS)
}

// Time complexity: O(len(n))
// Space complexity: O(len(n))
