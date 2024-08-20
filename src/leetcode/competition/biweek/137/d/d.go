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

todo
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

	pre := make([][]int, m)

	for j := range pre {
		pre[j] = make([]int, n)
	}

	for i, r := range board {
		for j, v := range r {
			boards[i] = append(boards[i], pair{v, j})
		}
		sort.Slice(boards[i], func(x, y int) bool {
			return boards[i][x].v > boards[i][y].v
		})
	}

	return 1
}

func maximumValueSum2(board [][]int) int64 {
	m := len(board)
	type pair struct{ x, j int }
	suf := make([][3]pair, m)
	p := [3]pair{} // 最大、次大、第三大
	for i := range p {
		p[i].x = math.MinInt
	}
	update := func(row []int) {
		for j, x := range row {
			if x > p[0].x {
				if p[0].j != j { // 如果相等，仅更新最大
					if p[1].j != j { // 如果相等，仅更新最大和次大
						p[2] = p[1]
					}
					p[1] = p[0]
				}
				p[0] = pair{x, j}
			} else if x > p[1].x && j != p[0].j {
				if p[1].j != j { // 如果相等，仅更新次大
					p[2] = p[1]
				}
				p[1] = pair{x, j}
			} else if x > p[2].x && j != p[0].j && j != p[1].j {
				p[2] = pair{x, j}
			}
		}
	}
	for i := m - 1; i > 1; i-- {
		update(board[i])
		suf[i] = p
	}

	ans := math.MinInt
	for i := range p {
		p[i].x = math.MinInt // 重置，计算 pre
	}
	for i, row := range board[:m-2] {
		update(row)
		for j, x := range board[i+1] { // 第二个车
			for _, p := range p { // 第一个车
				for _, q := range suf[i+2] { // 第三个车
					if p.j != j && q.j != j && p.j != q.j { // 没有同列的车
						ans = max(ans, p.x+x+q.x)
						break
					}
				}
			}
		}
	}
	return int64(ans)
}

//{{33,-5,10},
//{4,-33,-40},
//{-20,22,98}}

func main() {
	//fmt.Println(maximumValueSum([][]int{{-3, 1, 1, 1}, {-3, 1, -3, 1}, {-3, 2, 1, 1}}))
	fmt.Println(maximumValueSum([][]int{{33, -5, 10}, {4, -33, -40}, {-20, 22, 98}}))

	// 3*3*3 *500*3

	//27*27 = 125000
	fmt.Println(1 << 10)
}
