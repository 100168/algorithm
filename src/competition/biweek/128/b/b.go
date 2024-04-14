package main

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	cnt := 1
	pre := points[0][0]
	for _, v := range points {
		x := v[0]
		if x-pre > w {
			cnt++
			pre = x
		}

	}
	return cnt
}

func main() {
	println(minRectanglesToCoverPoints([][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1))
}
