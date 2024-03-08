package main

//有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
//wi 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
//
//
// 示例 3：
//
//
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 722 👎 0

import (
	"container/heap"
	"math"
)

type pair struct {
	x int
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

func networkDelayTime(times [][]int, n int, k int) int {

	type edge struct {
		to int
		d  int
	}

	g := make([][]edge, n+1)

	for _, t := range times {
		x, y, d := t[0], t[1], t[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}

	dis := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dis[i] = math.MaxInt
	}
	dis[k] = 0
	hp := &hp{{k, 0}}
	for hp.Len() > 0 {
		p := heap.Pop(hp).(pair)
		x := p.x
		if p.d > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newDis := p.d + e.d
			if newDis < dis[y] {
				dis[y] = newDis
				heap.Push(hp, pair{y, newDis})
			}
		}
	}
	ans := 0
	for i := range dis {
		ans = max(ans, dis[i])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func main() {
	println(networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
}
