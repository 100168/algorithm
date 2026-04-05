package main

import "slices"

func smallestDivisor(nums []int, threshold int) int {

	t := slices.Max(nums) + 1

	l, r := 1, t

	//循环不变量 r+1 <=threshold  l-1>threshold  ==>asn = r+1 == l
	for l <= r {
		m := (r + l) / 2
		if check(nums, m) <= threshold {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func check(nums []int, m int) int {

	sum := 0
	for _, v := range nums {
		sum += (v-1)/m + 1
	}
	return sum
}
