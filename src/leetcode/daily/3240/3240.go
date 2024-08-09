package main

import "fmt"

/*
*
给你一个 m x n 的二进制矩阵 grid 。

如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。

你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。

请你返回 最少 翻转次数，使得矩阵中 所有 行和列都是 回文的 ，且矩阵中 1 的数目可以被 4 整除 。
*/
func minFlips(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	//a[i][j] == a[m-i-1][j] == a[i][n-j-1][j]==a[m-i-1][n-j-1]

	cnt := 0
	for i, r := range grid[:m/2] {
		r2 := grid[m-1-i]

		for j := range r[:n/2] {
			cntOne := r[j] + r[n-j-1] + r2[j] + r2[n-j-1]
			cnt += min(cntOne, 4-cntOne)
		}

	}
	cntDiff := 0
	cntOne := 0
	if m&1 == 1 {
		r := grid[m/2]
		for j := range r[:n/2] {
			if r[j] != r[n-1-j] {
				cntDiff++
			} else if r[j] == 1 {
				cntOne += 2
			}
		}
	}
	if n&1 == 1 {
		for i, r1 := range grid[:m/2] {
			r2 := grid[m-i-1]
			if r1[n/2] != r2[n/2] {
				cntDiff++
			} else if r1[n/2] == 1 {
				cntOne += 2
			}
		}
	}
	if cntOne%4 != 0 && cntDiff == 0 {
		cnt += cntOne % 4
	} else {
		cnt += cntDiff
	}
	if m&1 == 1 && n&1 == 1 && grid[m/2][n/2] == 1 {
		cnt++
	}
	return cnt
}

//{{0,0,1},
//{0,0,1},
//{1,0,1},
//{1,0,0},
//{0,1,1}}

func main() {
	//fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	//fmt.Println(minFlips([][]int{{1}, {1}, {1}, {0}}))
	//fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))
	//fmt.Println(minFlips([][]int{{0}, {1}, {0}}))
	//fmt.Println(minFlips([][]int{{1}, {1}}))
	//fmt.Println(minFlips([][]int{{0}, {0}, {1}, {1}, {1}, {0}, {1}}))
	//fmt.Println(minFlips([][]int{{0, 0, 1}, {0, 0, 1}, {1, 0, 1}, {1, 0, 0}, {0, 1, 1}}))
	fmt.Println(minFlips([][]int{{1, 1, 0, 0, 1}, {0, 1, 1, 0, 1}, {0, 0, 1, 1, 1}, {1, 0, 1, 1, 0}, {1, 0, 1, 1, 0}, {0, 0, 1, 1, 1}, {0, 1, 1, 0, 1}, {1, 1, 0, 1, 0}}))
}
