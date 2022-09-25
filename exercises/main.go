package main

func partition(s []int, pivot int) (i int) {
	l, r := 0, len(s)-1
	p := s[pivot]
	for {
		for s[l] < p {
			l++
		}
		for s[r] >= p && r > 0 {
			r--
		}
		if l >= r && r > 0 {
			return r + 1
		} else if l >= r && r <= 0 {
			return r
		}
		s[l], s[r] = s[r], s[l]
	}
}

func main() {

}
