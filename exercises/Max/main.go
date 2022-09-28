package main

import "fmt"

func minMax(s []int, min, max int) {
	if len(s) == 0 {
		fmt.Println("min:", min, " max:", max)
		return
	}
	if s[0] > max {
		max = s[0]
	}
	if s[0] < min {
		min = s[0]
	}
	minMax(s[1:], min, max)
}
func main() {
	max := 0
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
	min := s[0]
	minMax(s, min, max)

}
