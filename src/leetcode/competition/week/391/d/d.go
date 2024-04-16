package main

import (
	"math"
	"sort"
)

func minimumDistance(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	n := len(points)
	l, r := 0, n-1
	ans := math.MaxInt

	for l < r {
		curL := points[l]
		curR := points[r]
		diff := abs(curL[0]-curR[0]) + abs(curL[1]-curR[1])
		ans = min(diff, ans)
		if curL[0] > curR[0] {
			r--
		} else {
			l++
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
