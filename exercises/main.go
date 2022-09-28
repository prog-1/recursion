package main

import "fmt"

func partitions(s []int, pivot int) (i int) {
	l, r := 0, len(s)-1
	p := s[pivot]
	for {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l >= r {
			return l
		}
		s[l], s[r] = s[r], s[l]
		r--
		// код писал я сам, по памяти с того что разбирали на уроке
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
	// time complexity = O(1)
	// space complaxity = O(1)
}
func powerof2(n int) bool {
	return (n & (n - 1)) == 0
	// time complexity = O(1)
	// space complaxity = O(1)
}
func sumdigits(n int, sum int) int {
	if n != 0 {
		i := n % 10
		sum += i
		return sumdigits(n/10, sum)
	}
	return sum
	//time complexity = O(n) or maybe O(log n) idk
	// space complaxity = O(n)
}

func maxnum(n []int, i int, max int, max2 int) (int, int) {
	if len(n) != i {
		if n[i] > max {
			max2 = max
			max = n[i]
		}
		i++
		return maxnum(n, i, max, max2)
	}
	return max, max2
	//time complexity = O(n)
	// space complaxity = O(n)
}

func main() {
	s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}
	i := partitions(s, 9)
	fmt.Println("Output: ", i, s)
	for i := 0; i <= 6; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println("\n", powerof2(256))
	fmt.Println(sumdigits(8999999991, 0))
	n := []int{10, 2, -3, 14}
	fmt.Println(maxnum(n, 0, 0, 0))
}
