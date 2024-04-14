package main

import (
	"container/heap"
	"math"
)

func minimumTime(n int, edges [][]int, disappear []int) []int {

	type edge struct{ to, d int }
	g := make([][]edge, n) // 邻接表
	for _, r := range edges {
		x, y, d := r[0], r[1], r[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	h := &hp{{}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		x := p.x
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] { // 尝试更新 x 的邻居的最短路
			y := e.to
			newDis := p.dis + e.d
			if newDis < dis[y] && newDis < disappear[y] {
				// 就目前来说，最短路必须经过 x
				dis[y] = newDis
				heap.Push(h, pair{newDis, y})
			}
		}
	}

	ans := make([]int, n)
	for i, v := range dis {
		if v > disappear[i] {
			ans[i] = -1
		} else {
			ans[i] = v
		}
	}

	return ans

}

type pair struct{ dis, x int }
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].dis < (*h)[j].dis }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)       { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
