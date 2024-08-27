package main

import (
	"container/heap"
	"fmt"
	"slices"
)

/**
给你一个整数数组 nums ，一个整数 k  和一个整数 multiplier 。

你需要对 nums 执行 k 次操作，每次操作中：

找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
将 x 替换为 x * multiplier 。
k 次操作以后，你需要将 nums 中每一个数值对 109 + 7 取余。

请你返回执行完 k 次乘运算以及取余运算之后，最终的 nums 数组。


*/

const mod = int(1e9 + 7)

func getFinalState(nums []int, k int, multiplier int) []int {

	if multiplier == 1 {
		return nums
	}
	n := len(nums)

	h := new(myHeap)
	for i, v := range nums {
		heap.Push(h, pair{i, v})
	}

	maxV := slices.Max(nums)

	for k > 0 {
		pop := heap.Pop(h).(pair)
		if pop.v == maxV || pop.v*multiplier > maxV {
			heap.Push(h, pop)
			break
		}
		if h.Len() > 0 {
			next := heap.Pop(h).(pair)
			for pop.v*multiplier < next.v && k > 0 {
				pop.v *= multiplier
				k--
			}
			heap.Push(h, next)
		}
		if k > 0 {
			pop.v *= multiplier
			k--
		}
		heap.Push(h, pop)
	}

	mul := k / n
	k = k % n

	for k > 0 {
		pop := heap.Pop(h).(pair)
		pop.v *= multiplier
		k--
		heap.Push(h, pop)
	}
	for h.Len() > 0 {
		pop := heap.Pop(h).(pair)
		nums[pop.index] = (pop.v % mod) * pow(multiplier, mul) % mod
	}
	return nums

}

func pow(a, b int) int {

	ans := 1
	for b > 0 {
		if b&1 == 1 {
			ans = ans * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return ans
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
	//fmt.Println(getFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
	//fmt.Println(getFinalState([]int{100000, 2000}, 2, 1000000))
	//fmt.Println(getFinalState([]int{1, 1, 4}, 4, 2))
	//fmt.Println(getFinalState([]int{3, 1, 3, 1, 3, 4, 5}, 5, 3))
	//fmt.Println(getFinalState([]int{66307295, 441787703, 589039035, 322281864}, 900900704, 641725))
	fmt.Println(getFinalState([]int{689009555, 813837455, 240010825, 967305450}, 804709161, 601392))
	fmt.Println(getFinalState([]int{889458628, 338743558, 875422936, 684907163, 233489834}, 246181588, 313380))
}
