package main

import (
	"fmt"
)

/*
*给你一个 n x n 的二进制网格 grid，每一次操作中，你可以选择网格的 相邻两行 进行交换。

一个符合要求的网格需要满足主对角线以上的格子全部都是 0 。

请你返回使网格满足要求的最少操作次数，如果无法使网格符合要求，请你返回 -1 。

主对角线指的是从 (1, 1) 到 (n, n) 的这些格子。

思路：贪心
*/
func minSwaps(grid [][]int) int {

	n := len(grid)
	ans := 0

	for r := 0; r < n; r++ {

		find := -1
		for i := r; i < n; i++ {
			cnt := 0
			for j := r + 1; j < n; j++ {
				cnt += grid[i][j]
			}
			if cnt == 0 {
				find = i
				break
			}
		}
		if find == -1 {
			return -1
		}
		ans += find - r
		for i := find; i > r; i-- {
			grid[i] = grid[i-1]
		}

	}
	return ans

}

func minSwaps2(grid [][]int) int {

	n := len(grid)
	ans := 0

	pos := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				pos[i] = j
			}
		}
	}
	for r := 0; r < n; r++ {

		find := -1
		for i := r; i < n; i++ {
			if pos[i] <= r {
				find = i
				break
			}
		}
		if find == -1 {
			return -1
		}
		ans += find - r
		for i := find; i > r; i-- {
			pos[i] = pos[i-1]
		}

	}
	return ans
}

func main() {
	//fmt.Println(minSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
	fmt.Println(minSwaps2([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
}
