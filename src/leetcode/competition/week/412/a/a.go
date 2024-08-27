package main

import "container/heap"

/*
*

给你一个整数数组 nums ，一个整数 k  和一个整数 multiplier 。

你需要对 nums 执行 k 次操作，每次操作中：

找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
将 x 替换为 x * multiplier 。
请你返回执行完 k 次乘运算之后，最终的 nums 数组。

示例 1：

输入：nums = [2,1,3,5,6], k = 5, multiplier = 2

输出：[8,4,6,5,6]
*/
func getFinalState(nums []int, k int, multiplier int) []int {

	h := new(myHeap)
	for i, v := range nums {

		heap.Push(h, pair{i, v})
	}

	for k > 0 {

		pop := heap.Pop(h).(pair)
		pop.v *= multiplier
		heap.Push(h, pop)
		k--
	}

	for h.Len() > 0 {

		pop := heap.Pop(h).(pair)
		nums[pop.index] = pop.v
	}
	return nums

}

type pair struct {
	index int
	v     int
}

type myHeap []pair

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].v < (*h)[j].v || (*h)[i].v == (*h)[j].v && (*h)[i].index < (*h)[j].index
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
	*h = append(*h, v.(pair))
}

func main() {

}
