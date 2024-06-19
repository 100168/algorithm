package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
*
给你一个下标从 1 开始、大小为 m x n 的整数矩阵 mat，你可以选择任一单元格作为 起始单元格 。

从起始单元格出发，你可以移动到 同一行或同一列 中的任何其他单元格，但前提是目标单元格的值 严格大于 当前单元格的值。

你可以多次重复这一过程，从一个单元格移动到另一个单元格，直到无法再进行任何移动。

请你找出从某个单元开始访问矩阵所能访问的 单元格的最大数量 。

返回一个表示可访问单元格最大数量的整数。

思路：
1.先排序
2.从小到到更新到当前位置的最大值
*/
func maxIncreasingCells(mat [][]int) int {

	m, n := len(mat), len(mat[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	type pair struct {
		i, j, v int
	}
	rows := make([][]pair, m)

	cols := make([][]pair, n)

	for i := range mat {
		for j := range mat[0] {
			rows[i] = append(rows[i], pair{i, j, mat[i][j]})
			cols[j] = append(cols[j], pair{i, j, mat[i][j]})
		}
	}

	for _, row := range rows {
		sort.Slice(row, func(i, j int) bool {
			return row[i].v < row[j].v
		})
	}
	for _, col := range cols {
		sort.Slice(col, func(i, j int) bool {
			return col[i].v < col[j].v
		})
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		cur := 0

		col := cols[j]
		row := rows[i]

		l, r := 0, len(col)-1

		for l <= r {
			m := (r + l) / 2
			if col[m].v > mat[i][j] {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		for k := l; k < m; k++ {
			cur = max(cur, dfs(col[k].i, col[k].j)+1)
		}

		l, r = 0, len(row)-1

		for l <= r {
			m := (r + l) / 2
			if row[m].v > mat[i][j] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		for k := l; k < n; k++ {
			cur = max(cur, dfs(row[k].i, row[k].j)+1)
		}

		memo[i][j] = cur
		return cur
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j)+1)
		}
	}
	return ans
}

func maxIncreasingCells2(mat [][]int) int {
	type pair struct{ x, y int }
	g := map[int][]pair{}
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j}) // 相同元素放在同一组，统计位置
		}
	}
	keys := make([]int, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range keys {
		pos := g[x]
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return slices.Max(rowMax)
}
func main() {

	fmt.Println(maxIncreasingCells([][]int{{3, 1, 6}, {-9, 5, 7}}))
}
