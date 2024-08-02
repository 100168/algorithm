package main

import (
	"fmt"
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

/**
给你一个正整数数组 nums，请你帮忙从该数组中找出能满足下面要求的 最长 前缀，并返回该前缀的长度：

从前缀中 恰好删除一个 元素后，剩下每个数字的出现次数都相同。
如果删除这个元素后没有剩余元素存在，仍可认为每个数字都具有相同的出现次数（也就是 0 次）。



示例 1：

输入：nums = [2,2,1,1,5,3,3,5]
输出：7
解释：对于长度为 7 的子数组 [2,2,1,1,5,3,3]，如果我们从中删去 nums[4] = 5，就可以得到 [2,2,1,1,3,3]，里面每个数字都出现了两次。

*/

func maxEqualFreq(nums []int) int {

	reb := redblacktree.New[int, int]()

	cntMap := make(map[int]int)
	ans := 0
	for i, v := range nums {
		cntMap[v]++
		put(reb, cntMap[v])
		remove(reb, cntMap[v]-1)

		left := reb.Left()
		right := reb.Right()

		if left.Value == 1 && right.Key*right.Value == i ||
			right.Value == 1 && right.Key-1 == left.Key ||
			right.Key == left.Key && (left.Key == 1 || left.Value == 1) {
			ans = i + 1
		}

	}
	return ans
}

func put(t *redblacktree.Tree[int, int], k int) {
	value, _ := t.Get(k)
	t.Put(k, value+1)
}

func remove(t *redblacktree.Tree[int, int], k int) {
	value, _ := t.Get(k)
	if value <= 1 {
		t.Remove(k)
	} else {
		t.Put(k, value-1)
	}
}

// 最大频率数
// 最大数的率*最大数
func maxEqualFreq2(nums []int) (ans int) {

	countMap := make(map[int]int)
	freqMap := make(map[int]int)

	maxCount := 0

	for i, v := range nums {
		if countMap[v] > 0 {
			freqMap[countMap[v]]--
		}
		countMap[v]++
		maxCount = max(maxCount, countMap[v])

		freqMap[countMap[v]]++
		// 分类讨论:
		//1. 每个数出现一次 maxCount ==1
		//2.去掉最小数满足 :freq[maxCount]*maxCount == i
		//3.去掉最大数满足: freqMap[maxCount-1]+1)*(maxCount-1)==i
		//4.只有一个数：maxCount = i+1

		if maxCount == 1 || maxCount == i+1 || freqMap[maxCount]*maxCount == i || (freqMap[maxCount-1]+1)*(maxCount-1) == i {
			ans = max(ans, i+1)
		}

	}
	return
}

func main() {
	//fmt.Println(maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}))
	//fmt.Println(maxEqualFreq([]int{1, 2}))
	//fmt.Println(maxEqualFreq([]int{1, 1}))
	//fmt.Println(maxEqualFreq2([]int{2, 2, 1, 1, 5, 3, 3, 5}))
	//fmt.Println(maxEqualFreq2([]int{1, 2}))
	//fmt.Println(maxEqualFreq2([]int{1, 1}))
	fmt.Println(maxEqualFreq2([]int{1, 1, 1, 2, 2, 2}))
}
