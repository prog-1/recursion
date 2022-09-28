package main

func cycle(n, sum int) int {
	sum += n % 10
	n /= 10
	if n == 0 {
		return sum
	}
	return cycle(n, sum)
}
func sumd(n int) int {
	if n < 0 {
		n *= -1
	}
	var sum int
	return cycle(n, sum)
}

func main() {
}
