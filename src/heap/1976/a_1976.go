package _976

import (
	"container/heap"
	"math"
)

func countPaths(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n) // 邻接表
	for _, r := range roads {
		x, y, d := r[0], r[1], r[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	f := make([]int, n)
	f[0] = 1
	h := &hp{{}}
	for {
		p := heap.Pop(h).(pair)
		x := p.x
		if x == n-1 {
			// 不可能找到比 dis[n-1] 更短，或者一样短的最短路了（注意本题边权都是正数）
			return f[n-1]
		}
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] { // 尝试更新 x 的邻居的最短路
			y := e.to
			newDis := p.dis + e.d
			if newDis < dis[y] {
				// 就目前来说，最短路必须经过 x
				dis[y] = newDis
				f[y] = f[x]
				heap.Push(h, pair{newDis, y})
			} else if newDis == dis[y] {
				// 和之前求的最短路一样长
				f[y] = (f[y] + f[x]) % 1_000_000_007
			}
		}
	}
}

type pair struct{ dis, x int }
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].dis < (*h)[j].dis }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)       { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
