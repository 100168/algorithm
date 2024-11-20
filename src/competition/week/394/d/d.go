package main

import (
	"container/heap"
	"math"
)

func findAnswer(n int, edges [][]int) []bool {
	m := len(edges)

	g := make([][]pair, n)
	for _, v := range edges {
		x, y, d := v[0], v[1], v[2]
		g[x] = append(g[x], pair{y, d})
		g[y] = append(g[y], pair{x, d})
	}

	f1 := make([]int, n)
	f2 := make([]int, n)

	for i := range f2 {
		f2[i] = math.MaxInt
	}
	for i := range f1 {
		f1[i] = math.MaxInt
	}
	f1[0] = 0
	h1 := &hp{{}}
	f2[n-1] = 0
	h2 := &hp{{n - 1, 0}}

	for h1.Len() > 0 {
		cur := heap.Pop(h1).(pair)
		x := cur.x
		if f1[x] < cur.d {
			continue
		}
		for _, e := range g[x] {
			newD := e.d + cur.d
			if newD < f1[e.x] {
				f1[e.x] = newD
				heap.Push(h1, pair{e.x, newD})
			}
		}
	}

	for h2.Len() > 0 {
		cur := heap.Pop(h2).(pair)
		x := cur.x
		if f2[x] < cur.d {
			continue
		}
		for _, e := range g[x] {
			newD := e.d + cur.d
			if newD < f2[e.x] {
				f2[e.x] = newD
				heap.Push(h2, pair{e.x, newD})
			}
		}
	}

	ans := make([]bool, m)
	for i, v := range edges {
		x, y, d := v[0], v[1], v[2]
		if f1[x]+f2[y]+d == f1[n-1] || f1[y]+f2[x]+d == f1[n-1] {
			ans[i] = true
		}
	}

	return ans
}

type pair struct{ x, d int }

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
func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
