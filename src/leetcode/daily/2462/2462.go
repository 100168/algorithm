package main

import (
	"container/heap"
	"math"
)

func totalCost(costs []int, k int, candidates int) int64 {

	lh := new(hp)
	rh := new(hp)

	n := len(costs)

	l := 0

	r := n - 1
	for ; l < candidates; l++ {
		heap.Push(lh, costs[l])
	}
	for ; r >= max(n-candidates, l); r-- {
		heap.Push(rh, costs[r])
	}

	ans := 0

	for k > 0 {

		leftMin := math.MaxInt
		rightMin := math.MaxInt
		if lh.Len() > 0 {
			leftMin = heap.Pop(lh).(int)
		}
		if rh.Len() > 0 {
			rightMin = heap.Pop(rh).(int)
		}
		if leftMin <= rightMin {
			heap.Push(rh, rightMin)
			if l <= r {
				heap.Push(lh, costs[l])
				l++
			}
		} else {
			if r >= l {
				heap.Push(rh, costs[r])
				r--
			}
			heap.Push(lh, leftMin)
		}
		k--
		ans += min(leftMin, rightMin)
	}

	return int64(ans)

}
func main() {

	println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
}

type hp []int

func (h *hp) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *hp) Swap(i, j int) {

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(int))
}
