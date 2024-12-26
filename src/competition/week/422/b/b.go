package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
*
有一个地窖，地窖中有 n x m 个房间，它们呈网格状排布。

给你一个大小为 n x m 的二维数组 moveTime ，
其中 moveTime[i][j] 表示在这个时刻 以后 你才可以 开始 往这个房间 移动 。
你在时刻 t = 0 时从房间 (0, 0) 出发，每次可以移动到 相邻 的一个房间。在 相邻 房间之间移动需要的时间为 1 秒。

请你返回到达房间 (n - 1, m - 1) 所需要的 最少 时间。

如果两个房间有一条公共边（可以是水平的也可以是竖直的），那么我们称这两个房间是 相邻 的。

示例 1：

输入：moveTime = [[0,4],[4,4]]

输出：6

解释：

需要花费的最少时间为 6 秒。

在时刻 t == 4 ，从房间 (0, 0) 移动到房间 (1, 0) ，花费 1 秒。
在时刻 t == 5 ，从房间 (1, 0) 移动到房间 (1, 1) ，花费 1 秒。
*/

type pair struct {
	x, y, d, v int
}

type hp []pair

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	return (*h)[i].d < (*h)[j].d
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() (v any) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func minTimeToReach(moveTime [][]int) int {

	m := len(moveTime)
	n := len(moveTime[0])
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	h := &hp{{0, 0, 0, 0}}

	ids := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t := i*n + j
			ids[t] = 0
		}
	}

	path := make([]int, m*n)

	for i := range path {
		path[i] = math.MaxInt
	}

	path[0] = 0
	for h.Len() > 0 {
		next := heap.Pop(h).(pair)
		if next.v < ids[next.x*n+next.y] {
			continue
		}
		for _, d := range dirs {
			nx, ny := next.x+d[0], next.y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				curD := max(next.d, moveTime[nx][ny]) + 1

				if curD <= path[nx*n+ny] {
					path[nx*n+ny] = curD
					ids[nx*n+ny]++
					heap.Push(h, pair{nx, ny, curD, ids[nx*n+ny]})
				}
			}
		}
	}
	return path[len(path)-1]
}

//{{0, 38, 10},
//{58, 29, 83}}

func main() {
	//fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}}))
	fmt.Println(minTimeToReach([][]int{{0, 38, 10}, {58, 29, 83}}))
}
