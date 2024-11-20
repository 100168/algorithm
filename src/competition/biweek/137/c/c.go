package main

import (
	"fmt"
	"math"
	"sort"
)

/**
给你一个 m x n 的二维整数数组 board ，它表示一个国际象棋棋盘，其中 board[i][j] 表示格子 (i, j) 的 价值 。

处于 同一行 或者 同一列 车会互相 攻击 。你需要在棋盘上放三个车，确保它们两两之间都 无法互相攻击 。

请你返回满足上述条件下，三个车所在格子 值 之和 最大 为多少。
*/

func maximumValueSum(board [][]int) int64 {

	m := len(board)
	n := len(board[0])
	f := make([]map[int]int, m)

	for i := range f {
		f[i] = make(map[int]int)
	}

	type pair struct {
		v     int
		index int
	}

	boards := make([][]pair, m)

	for i, r := range board {

		for j, v := range r {
			boards[i] = append(boards[i], pair{v, j})
		}
		sort.Slice(boards[i], func(x, y int) bool {
			return boards[i][x].v > boards[i][y].v
		})
	}
	// 8 个位
	var dfs func(int, int) int

	dfs = func(i int, maskC int) int {

		f1 := (1<<8 - 1) & maskC
		f2 := (1<<8 - 1) & (maskC >> 8)
		f3 := (1<<8 - 1) & (maskC >> 16)

		if f1 > 0 && f2 > 0 && f3 > 0 {
			return 0
		}

		if i < 0 {
			return math.MinInt / 2
		}
		if _, ok := f[i][maskC]; ok {
			return f[i][maskC]
		}

		cur := math.MinInt / 2
		for j := 0; j < min(n, 5); j++ {
			curIndex := boards[i][j].index
			v := boards[i][j].v
			if f1 == 0 {
				cur = max(cur, dfs(i-1, curIndex+1)+v)
			} else if f2 == 0 && curIndex != f1-1 {
				cur = max(cur, dfs(i-1, (curIndex+1)<<8|maskC)+v)
			} else if f3 == 0 && curIndex != f1-1 && curIndex != f2-1 {
				cur = max(cur, dfs(i-1, (curIndex+1)<<16|maskC)+v)
			}
		}
		cur = max(cur, dfs(i-1, maskC))

		//fmt.Println(i, cur, f1, f2, f3)
		f[i][maskC] = cur
		return cur
	}

	return int64(dfs(m-1, 0))
}

//{{33,-5,10},
//{4,-33,-40},
//{-20,22,98}}

func main() {
	//fmt.Println(maximumValueSum([][]int{{-3, 1, 1, 1}, {-3, 1, -3, 1}, {-3, 2, 1, 1}}))
	fmt.Println(maximumValueSum([][]int{{33, -5, 10}, {4, -33, -40}, {-20, 22, 98}}))
}
