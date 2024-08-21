package main

import (
	"slices"
	"sort"
)

/*
*
给你一个二进制矩阵 matrix ，它的大小为 m x n ，你可以将 matrix 中的 列 按任意顺序重新排列。

请你返回最优方案下将 matrix 重新排列后，全是 1 的子矩阵面积。
*/
func largestSubmatrix(matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])

	cnt := make([]int, n)

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				cnt[j] = 0
			} else {
				cnt[j]++
			}
		}
		cp := slices.Clone(cnt)
		sort.Ints(cp)
		for j := 0; j < n; j++ {
			ans = max(ans, cp[j]*(n-j))
		}
	}
	return ans

}
