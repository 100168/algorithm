package main

func maximumRemovals(s string, p string, removable []int) int {
	l, r := 0, len(removable)

	for l <= r {
		m := l + (r-l)/2
		if check(s, p, m, removable) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func check(s, p string, k int, removable []int) bool {

	marked := make(map[int]bool)
	for i := 0; i < k; i++ {
		marked[removable[i]] = true
	}
	p1, p2 := 0, 0
	for p1 < len(s) && p2 < len(p) {

		if s[p1] == p[p2] && !marked[p1] {
			p1++
			p2++
		} else {
			p1++
		}
	}
	return p2 == len(p)
}
