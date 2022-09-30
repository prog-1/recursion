package main

func maxNumLoop(s []int) int {
	max := s[0]
	for _, el := range s {
		if el > max {
			max = el
		}
	}

	return max
}

func maxNum(s []int) []int {
	// Getting max num index
	var maxTwo []int
	max := maxN(0, 1, s)
	maxTwo = append(maxTwo, s[max])

	// Removing max num from array
	tmp := s[:max]
	s = append(tmp, s[max+1:]...)

	// Getting max num one more time
	max = maxN(0, 1, s)
	maxTwo = append(maxTwo, s[max])

	return maxTwo

	// Time complexity O(N) -- ~2N times maxN will be called
	// Space complexity O(N) -- ~2N times we will add new return address,
	//								and also tmp and s allocation, but it's constant
}
func maxN(max int, i int, s []int) int {
	if s[i] > s[max] {
		max = i
	}
	if i < len(s)-1 {
		return maxN(max, i+1, s)
	} else {
		return max
	}
}
