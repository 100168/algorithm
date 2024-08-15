package main

import "math"

/**
给你一个由 正整数 组成、大小为 m x n 的矩阵 grid。
你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。
从值为 c1 的单元格移动到值为 c2 的单元格的得分为 c2 - c1 。


你可以从 任一 单元格开始，并且必须至少移动一次。

返回你能得到的 最大 总得分。
*/

func maxScore(grid [][]int) int {

	ans := math.MinInt

	n := len(grid[0])
	rowMin := make([]int, n)
	for i := range rowMin {
		rowMin[i] = math.MaxInt
	}
	for _, r := range grid {
		preMin := math.MaxInt
		for j, v := range r {
			ans = max(ans, v-min(preMin, rowMin[j]))
			rowMin[j] = min(rowMin[j], v)
			preMin = min(preMin, rowMin[j])
		}
	}
	return ans
}
