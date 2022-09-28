package main

func rearrange(s []int, pivot int) (i int) {
	l, r, p := 0, len(s)-1, s[pivot]
	for {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l >= r {
			return r
		}
		if s[l] != s[r] {
			s[l], s[r] = s[r], s[l]
		} else {
			r--
		}
	}
}
