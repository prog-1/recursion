package main

// Outputs first n fibonacci numbers
func fibonacci(n int) []int {
	s := []int{0, 1}
	fib(n, &s)
	return s
}

func fib(n int, s *[]int) {
	if len(*s) < n {
		(*s) = append((*s), (*s)[len(*s)-1]+(*s)[len(*s)-2])
		fib(n, s)
	}
	// Time complexity O(N) -- Because fib() gets called 1 time per element
	// Space complexity O(N) -- Due to new slice allocation happening every time append() is called
}
