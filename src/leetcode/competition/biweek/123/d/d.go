package main

import (
	"math"
	"sort"
)

/*
*
给你一个  n x 2 的二维数组 points ，它表示二维平面上的一些点坐标，其中 points[i] = [xi, yi] 。

计算点对 (A, B) 的数量，其中

A 在 B 的左上角，并且
它们形成的长方形中（或直线上）没有其它点（包括边界）。
返回数量。
*/
func numberOfPairs(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] > points[j][1]
	})

	ans := 0

	n := len(points)
	for i := range points {
		maxV := math.MinInt
		for j := i + 1; j < n; j++ {
			if points[j][1] > points[i][1] {
				continue
			}
			if points[j][1] > maxV {
				ans++
			}
			maxV = max(maxV, points[j][1])
		}
	}
	return ans
}
