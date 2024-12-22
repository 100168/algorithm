package main

import (
	"fmt"
)

/**
给你一个 m x n 的二进制矩阵 grid 。

如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。

你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。

请你返回 最少 翻转次数，使得矩阵中 所有 行和列都是 回文的 ，且矩阵中 1 的数目可以被 4 整除 。
*/

func minFlips(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	cntR := 0

	cntROne := 0
	cntRZero := 0
	cntRMZero := 0
	cntRMOne := 0

	for i := 0; i < m; i++ {
		l, r := 0, n-1
		for l < r {
			if grid[i][l] != grid[i][r] {
				cntR++
			} else {
				if grid[i][l] == 1 {
					cntROne += 2
				} else {
					cntRZero += 2
				}
			}
			l++
			r--
		}
		if l == r && grid[i][l] == 1 {
			cntROne++
			cntRMOne++
		} else if l == r && grid[i][l] == 0 {
			cntRMZero++
		}
	}

	cntC := 0
	cntCOne := 0
	cntCZero := 0
	cntCMZero := 0
	cntCMOne := 0
	for j := 0; j < n; j++ {
		l, r := 0, m-1
		for l < r {
			if grid[l][j] != grid[r][j] {
				cntC++
			} else {
				if grid[l][j] == 1 {
					cntCOne += 2
				} else {
					cntCZero += 2
				}
			}
			l++
			r--
		}
		if l == r && grid[l][j] == 1 {
			cntCOne++
			cntCMOne++
		} else if l == r && grid[l][j] == 0 {
			cntCMZero++
		}
	}

	rest1 := cntR % 4
	rest2 := cntC % 4

	f1 := cntR

	if (4-rest1)%2 != 0 || (4-rest1)%2 == 0 && cntR == 0 {
		diff := 4 - rest1
		if diff%2 == 0 {
			if cntCZero >= 2 {
				f1 += 2
			} else {
				f1 += 2
			}
		} else {
			if cntRMZero >= diff {
				f1 += diff
			} else {

			}
		}
	}
	f2 := cntC
	if 2*cntC < 4-rest2 {
		diff := 4 - rest2
		if cntCMZero >= diff {
			diff = 0
		} else {
			diff -= cntCMZero
		}

		if diff%2 == 1 || diff > cntCZero {
			f2 += cntCOne
		} else {
			f2 += 4 - rest2
		}
	}

	return min(f1, f2)

}

//{{1, 0, 0},
//{0, 1, 0},
//{0, 0, 1}}

func main() {
	//fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	//fmt.Println(minFlips([][]int{{1, 0, 1}, {0, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}
