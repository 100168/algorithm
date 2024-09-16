package main

import (
	"fmt"
	"slices"
)

/*
*、
给你一个长度为 n 的二维整数数组 coordinates 和一个整数 k ，其中 0 <= k < n 。

coordinates[i] = [xi, yi] 表示二维平面里一个点 (xi, yi) 。

如果一个点序列 (x1, y1), (x2, y2), (x3, y3), ..., (xm, ym) 满足以下条件，那么我们称它是一个长度为 m 的 上升序列 ：

对于所有满足 1 <= i < m 的 i 都有 xi < xi + 1 且 yi < yi + 1 。
对于所有 1 <= i <= m 的 i 对应的点 (xi, yi) 都在给定的坐标数组里。
请你返回包含坐标 coordinates[k] 的 最长上升路径 的长度。

示例 1：

输入：coordinates = [[3,1],[2,2],[4,1],[0,0],[5,3]], k = 1

输出：3

解释：

(0, 0) ，(2, 2) ，(5, 3) 是包含坐标 (2, 2) 的最长上升路径。

示例 2：

输入：coordinates = [[2,1],[7,0],[5,6]], k = 2

输出：2

解释：

(2, 1) ，(5, 6) 是包含坐标 (5, 6) 的最长上升路径。
*/
func maxPathLength(coordinates [][]int, k int) int {

	n := len(coordinates)

	ids := make([]int, n)

	for i := 0; i < n; i++ {
		ids[i] = i
	}

	slices.SortStableFunc(ids, func(i, j int) int {
		if coordinates[i][0] != coordinates[j][0] {
			return coordinates[i][0] - coordinates[j][0]
		}
		return coordinates[j][1] - coordinates[i][1]

	})
	start := -1

	for i := 0; i < n; i++ {

		if ids[i] == k {

			start = i
			break
		}
	}

	ans1 := 0

	var st []int

	startV := coordinates[ids[start]][1]
	visited := make([]bool, n)
	visited[start] = true
	for i := start + 1; i < n; i++ {
		v := coordinates[ids[i]][1]
		if v <= startV {
			continue
		}
		visited[i] = true
		l, r := 0, len(st)-1
		for l <= r {
			m := (r + l) / 2
			if st[m] >= v {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if l != len(st) {
			st[l] = v
		} else {
			st = append(st, v)
		}
		ans1 = max(ans1, len(st))
	}

	st = make([]int, 0)

	ans2 := 0
	for i := 0; i < start; i++ {
		v := coordinates[ids[i]][1]
		if v >= startV {
			continue
		}
		l, r := 0, len(st)-1
		for l <= r {
			m := (r + l) / 2
			if st[m] >= v {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if l != len(st) {
			st[l] = v
		} else {
			st = append(st, v)
		}
		ans2 = max(ans2, len(st))
	}

	return ans1 + ans2 + 1

}

func main() {
	fmt.Println(maxPathLength([][]int{{2, 1}, {7, 0}, {5, 6}}, 2))
	fmt.Println(maxPathLength([][]int{{3, 1}, {2, 2}, {4, 1}, {0, 0}, {5, 3}}, 1))
	fmt.Println(maxPathLength([][]int{{1, 3}, {8, 5}}, 0))
	fmt.Println(maxPathLength([][]int{{5, 0}, {5, 5}}, 1))
	fmt.Println(maxPathLength([][]int{{4, 7}, {6, 8}, {0, 3}, {6, 0}}, 0))
	fmt.Println(maxPathLength([][]int{{0, 1}, {0, 3}}, 0))
	fmt.Println(maxPathLength([][]int{{9, 5}, {1, 4}, {1, 7}, {9, 8}, {6, 4}, {6, 7}}, 3))
}
