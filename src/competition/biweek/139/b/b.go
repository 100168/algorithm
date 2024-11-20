package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
*
给你一个 m x n 的二进制矩形 grid 和一个整数 health 表示你的健康值。

你开始于矩形的左上角 (0, 0) ，你的目标是矩形的右下角 (m - 1, n - 1) 。

你可以在矩形中往上下左右相邻格子移动，但前提是你的健康值始终是 正数 。

对于格子 (i, j) ，如果 grid[i][j] = 1 ，那么这个格子视为 不安全 的，会使你的健康值减少 1 。

如果你可以到达最终的格子，请你返回 true ，否则返回 false 。

注意 ，当你在最终格子的时候，你的健康值也必须为 正数 。

示例 1：

输入：grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1

输出：true
*/

func findSafeWalk(grid [][]int, health int) bool {

	m := len(grid)
	n := len(grid[0])

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	dis := make([]int, m*n)

	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = grid[0][0]

	h := new(hp)
	heap.Push(h, pair{0, 0, grid[0][0]})

	for h.Len() > 0 {
		pop := heap.Pop(h).(pair)
		if pop.d > dis[pop.x*n+pop.y] {
			continue
		}
		for _, d := range dir {
			x, y := d[0]+pop.x, d[1]+pop.y
			if x >= 0 && x < m && y >= 0 && y < n {
				di := pop.d + grid[x][y]
				if di >= dis[x*n+y] {
					continue
				}
				dis[x*n+y] = di
				heap.Push(h, pair{x, y, di})
			}
		}
	}

	return dis[m*n-1] < health

}

type pair struct {
	x int
	y int
	d int
}

type hp []pair

func (h *hp) Less(i, j int) bool {
	return (*h)[i].d < (*h)[j].d
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]

	return v
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func main() {
	//fmt.Println(findSafeWalk([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}, 1))
	//fmt.Println(findSafeWalk([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 4))
	fmt.Println(findSafeWalk([][]int{{0, 1, 1, 0, 0, 0}, {1, 0, 1, 0, 0, 0}, {0, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}}, 3))
}
