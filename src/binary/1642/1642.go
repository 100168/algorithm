package main

import "sort"

func furthestBuilding(heights []int, bricks int, ladders int) int {

	n := len(heights)
	l, r := 0, n-1

	diff := make([]int, n)

	for i := 1; i < n; i++ {
		if heights[i] > heights[i-1] {
			diff[i] = heights[i] - heights[i-1]
		}
	}
	for l <= r {
		m := (r-l)/2 + l
		if check(diff, m, bricks, ladders) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func check(diff []int, t, bricks, ladders int) bool {

	cur := make([]int, 0)
	sum := 0
	for i := 0; i <= t; i++ {
		cur = append(cur, diff[i])
		sum += diff[i]
	}
	sort.Ints(cur)
	for i := 0; i <= t-ladders; i++ {
		bricks -= cur[i]
		if bricks < 0 {
			return false
		}
	}
	return true
}
