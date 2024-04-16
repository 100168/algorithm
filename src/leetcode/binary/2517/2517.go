package main

import (
	"slices"
	"sort"
)

func maximumTastiness(price []int, k int) int {

	l, r := 1, slices.Max(price)

	sort.Ints(price)
	for l <= r {
		m := (r + l) / 2

		if check(m, k, price) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func check(t, k int, price []int) bool {

	cnt := 1
	pre := 0
	for i, v := range price[1:] {
		if v-price[pre] >= t {
			cnt++
			pre = i + 1
		}

	}
	return cnt >= k
}
