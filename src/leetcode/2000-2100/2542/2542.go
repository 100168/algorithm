package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，两者长度都是 n ，再给你一个正整数 k 。
你必须从 nums1 中选一个长度为 k 的 子序列 对应的下标。

对于选择的下标 i0 ，i1 ，...， ik - 1 ，你的 分数 定义如下：

nums1 中下标对应元素求和，乘以 nums2 中下标对应元素的 最小值 。
用公式表示： (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]) 。
请你返回 最大 可能的分数。

一个数组的 子序列 下标是集合 {0, 1, ..., n-1} 中删除若干元素得到的剩余集合，也可以不删除任何元素。

示例 1：

输入：nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
输出：12
解释：
四个可能的子序列分数为：
- 选择下标 0 ，1 和 2 ，得到分数 (1+3+3) * min(2,1,3) = 7 。
- 选择下标 0 ，1 和 3 ，得到分数 (1+3+2) * min(2,1,4) = 6 。
- 选择下标 0 ，2 和 3 ，得到分数 (1+3+2) * min(2,3,4) = 12 。
- 选择下标 1 ，2 和 3 ，得到分数 (3+3+2) * min(1,3,4) = 8 。
所以最大分数为 12 。

思路：按nums2从大到小排序，并用堆来维护k个最大值和
*/
func maxScore(nums1 []int, nums2 []int, k int) int64 {

	type pair struct {
		n1 int
		n2 int
	}

	hp := new(myHeap)
	n := len(nums2)
	pairs := make([]pair, n)
	for i, v := range nums2 {
		pairs[i] = pair{nums1[i], v}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].n2 > pairs[j].n2
	})
	s := 0
	ans := 0
	for i, v := range pairs {
		if i < k {
			s += v.n1
			heap.Push(hp, v.n1)
		} else {

			if (*hp)[0] < v.n1 {
				s -= (*hp)[0]
				heap.Pop(hp)
				s += v.n1
				heap.Push(hp, v.n1)
			}
		}
		if i >= k-1 {
			ans = max(s*v.n2, ans)
		}
	}
	return int64(ans)
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
	fmt.Println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	fmt.Println(maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3))
}
