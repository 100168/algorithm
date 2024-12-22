package main

import (
	"fmt"
	"slices"
)

/*
*
给你一个长度为 n的整数数组nums和一个整数数组queries。

gcdPairs表示数组 nums中所有满足 0 <= i < j < n的数对 (nums[i], nums[j]) 的
最大公约数升序排列构成的数组。

对于每个查询queries[i]，你需要找到gcdPairs中下标为queries[i]的元素。

请你返回一个整数数组answer，其中answer[i]是gcdPairs[queries[i]]的值。

gcd(a, b)表示 a和 b的 最大公约数。

示例 1：

输入：nums = [2,3,4], queries = [0,2,2]

输出：[1,2,2]

解释：

gcdPairs = [gcd(nums[0], nums[1]), gcd(nums[0], nums[2]), gcd(nums[1], nums[2])] = [1, 2, 1].

升序排序后得到gcdPairs = [1, 1, 2]。

所以答案为[gcdPairs[queries[0]], gcdPairs[queries[1]], gcdPairs[queries[2]]] = [1, 2, 2]。

思路：容斥+前缀和+二分


*/

func gcdValues(nums []int, queries []int64) []int {
	mx := slices.Max(nums)
	cntX := make([]int, mx+1)
	for _, x := range nums {
		cntX[x]++
	}

	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		c := 0
		for j := i; j <= mx; j += i {
			c += cntX[j]
			cntGcd[i] -= cntGcd[j] // gcd 是 2i,3i,4i,... 的数对不能统计进来
		}
		cntGcd[i] += c * (c - 1) / 2 // c 个数选 2 个，组成 c*(c-1)/2 个数对
	}

	for i := 2; i <= mx; i++ {
		cntGcd[i] += cntGcd[i-1] // 原地求前缀和
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = find(cntGcd, int(q)+1)
	}
	return ans
}
func find(nums []int, t int) int {
	// find >=t most left
	l, r := 0, len(nums)
	for l <= r {
		m := (r + l) / 2
		if nums[m] < t {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func main() {
	fmt.Println(gcdValues([]int{2, 3, 4}, []int64{0, 2, 2}))
}
