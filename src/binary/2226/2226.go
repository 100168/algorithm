package main

import "slices"

func maximumCandies(candies []int, k int64) int {

	l, r := 1, slices.Max(candies)

	for l <= r {
		m := (r + l) / 2

		//循环不变量 l-1 >=k ,r+1<k  ==> ans = r
		if check(candies, m) >= k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func check(candies []int, target int) int64 {

	sum := int64(0)
	for _, v := range candies {
		sum += int64(v / target)
	}
	return sum
}
