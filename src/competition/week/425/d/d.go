package main

import (
	"container/heap"
	"fmt"
)

func maximizeSumOfWeights(edges [][]int, k int) int64 {

	type pair struct {
		x int
		d int
	}

	n := len(edges) + 1

	g := make([][]pair, n)

	for _, v := range edges {

		x, y, d := v[0], v[1], v[2]

		g[x] = append(g[x], pair{y, d})
		g[y] = append(g[y], pair{x, d})
	}

	var dfs func(int, int) int

	dfs = func(x int, p int) int {

		h := &myHeap{}

		s := 0

		for _, v := range g[x] {
			if v.x == p {
				continue
			}
			next := dfs(v.x, x) + v.d

			heap.Push(h, next)
			s += next

			if h.Len() > k {
				s -= heap.Pop(h).(int)
			}
		}
		return s
	}

	return int64(dfs(0, -1))
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(int))
}
func main() {
	fmt.Println()
}
