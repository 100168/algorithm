package main

import (
	"container/heap"
	"fmt"
)

/*
*
汽车从起点出发驶向目的地，该目的地位于出发位置东面 target 英里处。

沿途有加油站，用数组 stations 表示。其中 stations[i] = [positioni, fueli]
表示第 i 个加油站位于出发位置东面 positioni 英里处，并且有 fueli 升汽油。

假设汽车油箱的容量是无限的，其中最初有 startFuel 升燃料。
它每行驶 1 英里就会用掉 1 升汽油。当汽车到达加油站时，它可能停下来加油，将所有汽油从加油站转移到汽车中。

为了到达目的地，汽车所必要的最低加油次数是多少？如果无法到达目的地，则返回 -1 。

注意：如果汽车到达加油站时剩余燃料为 0，它仍然可以在那里加油。如果汽车到达目的地时剩余燃料为 0，仍然认为它已经到达目的地。

示例 1：

输入：target = 1, startFuel = 1, stations = []
输出：0
解释：可以在不加油的情况下到达目的地。
示例 2：

输入：target = 100, startFuel = 1, stations = [[10,100]]
输出：-1
解释：无法抵达目的地，甚至无法到达第一个加油站。
*/
func minRefuelStops(target int, startFuel int, stations [][]int) int {

	ans := 0
	h := new(myHeap)

	stations = append(stations, []int{target, 0})

	pre := 0
	for _, v := range stations {
		startFuel -= v[0] - pre
		if startFuel <= 0 {
			for h.Len() > 0 && startFuel < 0 {
				pop := heap.Pop(h).(int)
				startFuel += pop
				ans++
			}
			if startFuel < 0 {
				return -1
			}
		}
		heap.Push(h, v[1])
		pre = v[0]
	}
	return ans
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
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
	fmt.Println(minRefuelStops(1, 1, [][]int{}))
	fmt.Println(minRefuelStops(100, 1, [][]int{{10, 100}}))
}
