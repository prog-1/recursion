package main

// func sumDigitsLoop(n int) int {
// 	var sum int
// 	for ; n != 0; n /= 10 {
// 		sum += n % 10
// 	}
// 	return sum
// }

func sumDigits(n int) int {
	if n != 0 {
		return n%10 + sumDigits(n/10)
	} else {
		return 0
	}

	// Time complexity O(N) -- One function call for each digit
	// Space complexity O(N) -- One return address for each digit
}
