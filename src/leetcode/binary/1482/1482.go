package main

import "slices"

func minDays(bloomDay []int, t int, k int) int {

	if len(bloomDay)/k < t {
		return -1
	}
	maxValue := slices.Max(bloomDay)
	l, r := 1, maxValue
	for l <= r {
		m := (r-l)/2 + l
		if check(bloomDay, m, k) >= t {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}

func check(bloomDay []int, t, k int) int {

	sum := 0
	cnt := 0
	for _, v := range bloomDay {

		if v <= t {
			cnt++
			if cnt == k {
				cnt = 0
				sum++
			}
		} else {
			cnt = 0
		}

	}
	return sum

}
