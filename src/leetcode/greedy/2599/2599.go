package main

import "container/heap"

/*
*
给定一个 下标从0开始 的整数数组 nums 。你可以任意多次执行以下操作：

从 nums 中选择任意一个元素，并将其放到 nums 的末尾。
nums 的前缀和数组是一个与 nums 长度相同的数组 prefix ，其中 prefix[i] 是所有整数 nums[j]（其中 j 在包括区间 [0，i] 内）的总和。

返回使前缀和数组不包含负整数的最小操作次数。测试用例的生成方式保证始终可以使前缀和数组非负。

示例 1 ：

输入：nums = [2,3,-5,4]
输出：0
解释：我们不需要执行任何操作。
给定数组为 [2, 3, -5, 4]，它的前缀和数组是 [2, 5, 0, 4]。
示例 2 ：

输入：nums = [3,-5,-2,6]
输出：1
解释：我们可以对索引为1的元素执行一次操作。
操作后的数组为 [3, -2, 6, -5]，它的前缀和数组是 [3, 1, 7, 2]。
*/
func makePrefSumNonNegative(nums []int) int {

	s := 0

	h := new(myHeap)

	ops := 0
	for _, v := range nums {

		s += v
		if v < 0 {
			heap.Push(h, v)
		}
		if s < 0 {
			s -= heap.Pop(h).(int)
			ops++
			nums = append(nums, v)
		}
	}
	return ops
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
