package main

import "slices"

/**
给你一个由正整数构成的二维矩阵 grid。

你需要从矩阵中选择 一个或多个 单元格，选中的单元格应满足以下条件：

所选单元格中的任意两个单元格都不会处于矩阵的 同一行。
所选单元格的值 互不相同。
你的得分为所选单元格值的总和。

返回你能获得的 最大 得分。


思路
1。值域状压dp
*/

func maxScore(grid [][]int) int {

	n := len(grid)

	mx := grid[0][0]

	indexMap := make(map[int][]int)

	for i, r := range grid {

		r = slices.Compact(r)
		for _, v := range r {
			mx = max(mx, v)
			indexMap[v] = append(indexMap[v], i)
		}
	}

	f := make([][]int, mx+1)

	for i := range f {
		f[i] = make([]int, 1<<n)

		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i int, mask int) int {

		if i < 0 {
			return 0
		}

		if f[i][mask] != -1 {
			return f[i][mask]
		}

		cur := 0
		for _, v := range indexMap[i] {
			if 1<<v&mask == 0 {
				cur = max(cur, dfs(i-1, 1<<v|mask)+i)
			}
		}

		cur = max(cur, dfs(i-1, mask))
		f[i][mask] = cur
		return cur
	}
	return dfs(mx, 0)
}
