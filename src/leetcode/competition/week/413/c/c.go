package main

import (
	"fmt"
	"sort"
)

/**
给你一个由正整数构成的二维矩阵 grid。

你需要从矩阵中选择 一个或多个 单元格，选中的单元格应满足以下条件：

所选单元格中的任意两个单元格都不会处于矩阵的 同一行。
所选单元格的值 互不相同。
你的得分为所选单元格值的总和。

返回你能获得的 最大 得分。
*/

func maxScore(grid [][]int) int {

	n := len(grid)

	interMap := make([]int, 101)
	for _, v := range grid {
		sort.Slice(v, func(i, j int) bool { return v[i] > v[j] })

		pre := -1
		for _, x := range v {
			if x != pre {
				interMap[x]++
			}
			pre = x
		}
	}

	for _, v := range grid {

		for _, x := range v {

			if interMap[x] > 0 {
			}
		}
	}

}

func maxScore2(grid [][]int) int {

	n := len(grid)

	type pair struct{ x, y int }

	f := make([]map[pair]int, n)

	for i := range f {

		f[i] = make(map[pair]int)
	}

	var dfs func(int, pair) int

	dfs = func(i int, mask pair) int {

		if i < 0 {
			return 0
		}
		if _, ok := f[i][mask]; ok {
			return f[i][mask]
		}

		cur := 0
		for _, v := range grid[i] {
			x, y := mask.x, mask.y

			if v > 50 {
				if 1<<(v-50)&x != 0 {
					continue
				} else {
					cur = max(cur, dfs(i-1, pair{x | 1<<(v-50), y})+v)
				}
			} else {
				if 1<<v&y != 0 {
					continue
				} else {
					cur = max(cur, dfs(i-1, pair{x, y | 1<<v})+v)
				}
			}
		}
		cur = max(cur, dfs(i-1, mask))
		f[i][mask] = cur
		return cur
	}
	return dfs(n-1, pair{0, 0})
}
func main() {
	fmt.Println(maxScore([][]int{{4}, {13}, {13}}))
}
