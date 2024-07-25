package main

import (
	"container/heap"
	"fmt"
	"math"
)

/**
给你一个整数数组 nums 和一个整数 k ，请你返回 非空 子序列元素和的最大值，
子序列需要满足：子序列中每两个 相邻 的整数 nums[i] 和 nums[j] ，
它们在原数组中的下标 i 和 j 满足 i < j 且 j - i <= k 。

数组的子序列定义为：将数组中的若干个数字删除（可以删除 0 个数字），剩下的数字按照原本的顺序排布。

示例 1：

输入：nums = [10,2,-10,5,20], k = 2
输出：37
解释：子序列为 [10, 2, 5, 20] 。
示例 2：

输入：nums = [-1,-2,-3], k = 1
输出：-1
解释：子序列必须是非空的，所以我们选择最大的数字。
*/

func constrainedSubsetSum(nums []int, k int) int {
	ans := math.MinInt
	hp := new(myHeap)
	for i, v := range nums {
		pop := pair{-100, 0}
		for hp.Len() > 0 {
			pop = heap.Pop(hp).(pair)
			if i-pop.index > k {
				continue
			}
			break
		}
		cur := pair{i, max(pop.v+v, v)}
		ans = max(ans, cur.v)
		heap.Push(hp, cur)
		heap.Push(hp, pop)
	}
	return ans
}

func constrainedSubsetSum2(nums []int, k int) int {
	ans := math.MinInt
	type pair struct {
		index int
		v     int
	}
	queue := make([]pair, 0)
	for i, v := range nums {

		var pop pair
		for len(queue) > 0 && (i-queue[0].index > k) {
			queue = queue[1:len(queue)]
		}
		if len(queue) > 0 {
			pop = queue[0]
		}
		cur := pair{i, max(v, v+pop.v)}
		for len(queue) > 0 && (i-queue[len(queue)-1].index > k || queue[len(queue)-1].v <= cur.v) {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, cur)
		ans = max(ans, cur.v)
	}
	return ans
}

type pair struct {
	index int
	v     int
}
type myHeap []pair

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].v > (*h)[j].v
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
	fmt.Println(constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2))
	fmt.Println(constrainedSubsetSum([]int{-5266, 4019, 7336, -3681, -5767}, 2))
	fmt.Println(constrainedSubsetSum2([]int{-5266, 4019, 7336, -3681, -5767}, 2))
}
