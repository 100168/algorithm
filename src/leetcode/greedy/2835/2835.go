package main

import (
	"container/heap"
	"fmt"
	"math/bits"
)

/*
*给你一个下标从 0 开始的数组 nums ，它包含 非负 整数，且全部为 2 的幂，同时给你一个整数 target 。

一次操作中，你必须对数组做以下修改：

选择数组中一个元素 nums[i] ，满足 nums[i] > 1 。
将 nums[i] 从数组中删除。
在 nums 的 末尾 添加 两个 数，值都为 nums[i] / 2 。
你的目标是让 nums 的一个 子序列 的元素和等于 target ，请你返回达成这一目标的 最少操作次数 。如果无法得到这样的子序列，请你返回 -1 。

数组中一个 子序列 是通过删除原数组中一些元素，并且不改变剩余元素顺序得到的剩余数组。

示例 1：

输入：nums = [1,2,8], target = 7
输出：1
解释：第一次操作中，我们选择元素 nums[2] 。数组变为 nums = [1,2,4,4] 。
这时候，nums 包含子序列 [1,2,4] ，和为 7 。
无法通过更少的操作得到和为 7 的子序列。

示例 2：

输入：nums = [1,32,1,2], target = 12
输出：2
解释：第一次操作中，我们选择元素 nums[1] 。数组变为 nums = [1,1,2,16,16] 。
第二次操作中，我们选择元素 nums[3] 。数组变为 nums = [1,1,2,16,8,8] 。
这时候，nums 包含子序列 [1,1,2,8] ，和为 12 。
无法通过更少的操作得到和为 12 的子序列。
1，2，3，4
*/
func minOperations(nums []int, target int) int {

	s := 0

	for _, v := range nums {
		s += v
	}

	if s < target {
		return -1
	}

	s = 0
	ops := 0

	hp := new(myHeap)
	for _, v := range nums {
		heap.Push(hp, v)
	}
	mostRight := bits.TrailingZeros(uint(target))
	s = 0
	for target > 0 {
		for s < 1<<mostRight && (*hp)[0] < 1<<mostRight {
			s += heap.Pop(hp).(int)
		}
		if s >= 1<<mostRight {
			s -= 1 << mostRight
		} else {
			pop := heap.Pop(hp).(int)
			index := bits.TrailingZeros(uint(pop))
			ops += index - mostRight
			for index > mostRight {
				heap.Push(hp, 1<<(index-1))
				index--
			}
		}
		target &= target - 1

		mostRight = bits.TrailingZeros(uint(target))
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

func minOperations2(nums []int, target int) (ans int) {
	s := 0
	cnt := [31]int{}
	for _, v := range nums {
		s += v
		cnt[bits.TrailingZeros(uint(v))]++
	}
	if s < target {
		return -1
	}
	s = 0
	for i := 0; 1<<i <= target; {
		s += cnt[i] << i
		mask := 1<<(i+1) - 1
		if s >= target&mask {
			i++
			continue
		}
		ans++
		for i++; cnt[i] == 0; i++ {
			ans++
		}
	}
	return
}

func main() {
	fmt.Println(minOperations([]int{1, 2, 8}, 7))
	fmt.Println(minOperations([]int{1024}, 400))
	fmt.Println(1 << 30)
}
