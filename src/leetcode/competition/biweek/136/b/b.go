package main

import "fmt"

/**
给你一个 m x n 的二进制矩阵 grid 。

如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。

你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。

请你返回 最少 翻转次数，使得矩阵 要么 所有行是 回文的 ，要么所有列是 回文的 。
*/

func minFlips(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	cntR := 0

	for i := 0; i < m; i++ {

		l, r := 0, n-1
		for l < r {
			if grid[i][l] != grid[i][r] {
				cntR++
			}
			l++
			r--
		}
	}

	cntC := 0
	for j := 0; j < n; j++ {
		l, r := 0, m-1
		for l < r {
			if grid[l][j] != grid[r][j] {
				cntC++
			}
			l++
			r--
		}
	}
	return min(cntR, cntC)

}

// [[0,0],
// [1,0],
// [1,1],
// [1,0]]
func main() {
	fmt.Println(minFlips([][]int{{0, 0}, {1, 0}, {1, 1}, {1, 0}}))
}
