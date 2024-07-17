package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
给你一个数组 events，其中 events[i] = [startDayi, endDayi] ，表示会议 i 开始于 startDayi ，结束于 endDayi 。

你可以在满足 startDayi <= d <= endDayi 中的任意一天 d 参加会议 i 。在任意一天 d 中只能参加一场会议。

请你返回你可以参加的 最大 会议数目。

堆+贪心：
1.
*/
func maxEvents(events [][]int) int {

	hp := &myHeap{}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	n := len(events)

	//  2,2,3,3,4,4  1,4

	cnt := 0
	//开始时间
	start := 1
	//当前位置
	i := 0

	for hp.Len() > 0 || i < n {
		//所有开始时间小于start的加入堆中（加入的是—end）
		for ; i < n && events[i][0] <= start; i++ {
			heap.Push(hp, events[i][1])
		}
		//把所有end小于start的弹出堆
		for hp.Len() > 0 && (*hp)[0] < start {
			heap.Pop(hp)
		}
		if hp.Len() > 0 {
			heap.Pop(hp)
			cnt++
		}
		start++
	}

	return cnt

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
	fmt.Println(maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}}))
	fmt.Println(maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}}))
}
