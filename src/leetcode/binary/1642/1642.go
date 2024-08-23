package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
给你一个整数数组 heights ，表示建筑物的高度。另有一些砖块 bricks 和梯子 ladders 。

你从建筑物 0 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。

当从建筑物 i 移动到建筑物 i+1（下标 从 0 开始 ）时：

如果当前建筑物的高度 大于或等于 下一建筑物的高度，则不需要梯子或砖块
如果当前建筑的高度 小于 下一个建筑的高度，您可以使用 一架梯子 或 (h[i+1] - h[i]) 个砖块
如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标 从 0 开始 ）。
*/
func furthestBuilding(heights []int, bricks int, ladders int) int {

	n := len(heights)
	l, r := 0, n-1

	diff := make([]int, n)

	for i := 1; i < n; i++ {
		if heights[i] > heights[i-1] {
			diff[i] = heights[i] - heights[i-1]
		}
	}
	for l <= r {
		m := (r-l)/2 + l
		if check(diff, m, bricks, ladders) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func check(diff []int, t, bricks, ladders int) bool {

	cur := make([]int, 0)
	sum := 0
	for i := 0; i <= t; i++ {
		cur = append(cur, diff[i])
		sum += diff[i]
	}
	sort.Ints(cur)
	for i := 0; i <= t-ladders; i++ {
		bricks -= cur[i]
		if bricks < 0 {
			return false
		}
	}
	return true
}

func furthestBuilding2(heights []int, bricks int, ladders int) int {

	h := new(myHeap)

	for i, v := range heights[:len(heights)-1] {

		i++
		next := heights[i]

		diff := next - v

		if diff > 0 {
			ladders--
			heap.Push(h, diff)
		}

		if ladders < 0 {

			pop := heap.Pop(h).(int)
			bricks -= pop
			if bricks < 0 {
				return i - 1
			}
			ladders++
		}
	}
	return len(heights) - 1
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
	//fmt.Println(furthestBuilding2([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	//fmt.Println(furthestBuilding2([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(furthestBuilding2([]int{14, 3, 19, 3}, 17, 0))
}
