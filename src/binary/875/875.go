package main

import "math"

func minEatingSpeed(piles []int, h int) int {

	l, r := 1, math.MaxInt

	for l <= r {
		m := (r-l)/2 + l
		if check(piles, m) <= h {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func check(piles []int, k int) int {

	ans := 0

	for _, v := range piles {
		ans += (v-1)/k + 1
	}
	return ans
}
