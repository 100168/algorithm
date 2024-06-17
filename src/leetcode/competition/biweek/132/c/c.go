package main

import (
	"fmt"
)

/*
*
给你一个整数数组 nums 和一个 非负 整数 k 。如果一个整数序列 seq 满足在范围下标范围 [0, seq.length - 2]
中存在 不超过 k 个下标 i 满足 seq[i] != seq[i + 1] ，那么我们称这个整数序列为 好 序列。

请你返回 nums 中 好
子序列

	的最长长度
*/

// Segment 表示一个由相同数字组成的段
type pair struct {
	Length int
	Value  int
}

// longestSubsequence 找到满足条件的“好”子序列的最长长度
func maximumLength(nums []int, k int) int {
	cntMap := make(map[int][]int)
	for i, v := range nums {
		cntMap[v] = append(cntMap[v], i)
	}
	ans := 0
	for _, v := range cntMap {
		l := 0
		cnt := 0
		for j, index := range v {
			if j > 0 && index != v[j-1]+1 {
				cnt += index - v[j-1] - 1
			}
			if cnt > k {
				for l+1 <= j && v[l] == v[l+1] {
					l++
				}
				l++
				cnt -= v[l] - v[max(l-1, 0)] - 1

			}
			ans = max(ans, j-l+1+cnt)
		}

	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1}
	k := 0
	maxLength := maximumLength(nums, k)
	fmt.Println("The longest good subsequence length is:", maxLength)
}
