package main

import (
	"container/heap"
	"fmt"
)

/*
*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

如果子数组中所有元素都相等，则认为子数组是一个 等值子数组 。注意，空数组是 等值子数组 。

从 nums 中删除最多 k 个元素后，返回可能的最长等值子数组的长度。

子数组 是数组中一个连续且可能为空的元素序列。

输入：nums = [1,3,2,3,1,3], k = 3
输出：3
解释：最优的方案是删除下标 2 和下标 4 的元素。
删除后，nums 等于 [1, 3, 3, 3] 。
最长等值子数组从 i = 1 开始到 j = 3 结束，长度等于 3 。
可以证明无法创建更长的等值子数组。
*/

/*
*
可以优化方法二中的堆
*/
func longestEqualSubarray(nums []int, k int) int {
	n := len(nums)
	ans := 0
	cnt := make(map[int]int)
	for i, j := 0, 0; j < n; j++ {
		cnt[nums[j]]++
		//不需要用堆来维护最值，只需要维护窗口的大小，并且每次来最新值更新
		for j-i+1 > cnt[nums[i]]+k {
			cnt[nums[i]]--
			i++
		}
		ans = max(ans, cnt[nums[i]])
	}
	return ans
}

func longestEqualSubarray3(nums []int, k int) (ans int) {
	posLists := make(map[int][]int)
	for i, x := range nums {
		posLists[x] = append(posLists[x], i)
	}

	for _, pos := range posLists {
		if len(pos) <= ans {
			continue // 无法让 ans 变得更大
		}
		left := 0
		for right, p := range pos {
			for p-pos[left]-(right-left) > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
		}
	}
	return
}

func longestEqualSubarray2(nums []int, k int) int {

	n := len(nums)
	cnt := make(map[int]int)
	ans := 0
	hp := new(myHeap)
	l := 0
	for i := 0; i < n; i++ {
		cnt[nums[i]]++
		heap.Push(hp, pair{cnt[nums[i]], nums[i]})
		pop := heap.Pop(hp).(pair)
		for pop.cnt != cnt[pop.val] {
			pop = heap.Pop(hp).(pair)
		}

		for pop.cnt+k < i-l+1 {
			cnt[nums[l]]--
			heap.Push(hp, pair{cnt[nums[l]], nums[l]})
			if nums[l] == pop.val {
				for pop.cnt != cnt[pop.val] {
					pop = heap.Pop(hp).(pair)
				}
			}
			l++
		}
		ans = max(ans, pop.cnt)
		heap.Push(hp, pop)

	}
	return ans
}

type pair struct {
	cnt int
	val int
}
type myHeap []pair

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].cnt > (*h)[j].cnt
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

	//fmt.Println(longestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(longestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
}
