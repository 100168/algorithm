package main

import (
	"slices"
	"sort"
)

func maxDistance(position []int, m int) int {

	l, r := 1, slices.Max(position)
	sort.Ints(position)
	for l <= r {
		mid := (r + l) / 2
		if check(mid, m, position) {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return r
}

func check(t, m int, position []int) bool {
	pre := 0
	m--
	for i, v := range position[1:] {
		if v-position[pre] >= t {
			m--
			pre = i + 1
		}
		if m == 0 {
			return true
		}

	}
	return false
}

func main() {
	println(maxDistance([]int{1, 2, 3, 45, 6, 8, 9, 100, 102, 103, 150}, 5))
}
