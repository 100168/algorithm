package main

import "fmt"

/**
给你一个二维字符矩阵 grid，其中 grid[i][j] 可能是 'X'、'Y' 或 '.'，返回满足以下条件的
子矩阵
数量：

包含 grid[0][0]
'X' 和 'Y' 的频数相等。
至少包含一个 'X'。
*/

func numberOfSubmatrices(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])

	sumX := make([][]int, m+1)
	for i := range sumX {
		sumX[i] = make([]int, n+1)
	}

	sumY := make([][]int, m+1)
	for i := range sumY {
		sumY[i] = make([]int, n+1)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				sumX[i+1][j+1] = sumX[i+1][j] + sumX[i][j+1] - sumX[i][j] + 1
			} else {
				sumX[i+1][j+1] = sumX[i+1][j] + sumX[i][j+1] - sumX[i][j]
			}
			if grid[i][j] == 'Y' {
				sumY[i+1][j+1] = sumY[i+1][j] + sumY[i][j+1] - sumY[i][j] + 1
			} else {
				sumY[i+1][j+1] = sumY[i+1][j] + sumY[i][j+1] - sumY[i][j]
			}
			if sumX[i+1][j+1] > 0 && sumX[i+1][j+1] == sumY[i+1][j+1] {
				ans++
			}

		}
	}
	return ans

}

func main() {
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}))
}
