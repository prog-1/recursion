package main

func cycle(n, cnt int, s []int) []int {
	s = append(s, s[len(s)-1]+s[len(s)-2])
	cnt++
	if cnt != n {
		return cycle(n, cnt, s)
	}
	return s
}
func fib(n int) []int {
	if n < 1 {
		return nil
	} else if n == 0 {
		return nil
	} else if n == 1 {
		return []int{0, 1}
	}
	cnt := 1
	var s = []int{0, 1}
	return cycle(n, cnt, s)
}
func main() {

}
