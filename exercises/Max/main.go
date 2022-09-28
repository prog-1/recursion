package main

import "fmt"

func cycle(s []int, min, max int) (int, int) {
	if len(s) == 0 {
		fmt.Println("min:", min, " max:", max)
		return min, max
	}
	if s[0] > max {
		max = s[0]
	}
	if s[0] < min {
		min = s[0]
	}
	return cycle(s[1:], min, max)
}
func minMax(s []int) (min, max int) {
	if len(s) != 0 {
		min = s[0]
		return cycle(s, min, max)
	} else {
		return 0, 0
	}

}
func main() {

}
