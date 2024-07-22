package main

import "sort"

/*
*
给你一个由正整数组成的数组 nums 和一个 正 整数 k 。

如果 nums 的子集中，任意两个整数的绝对差均不等于 k ，则认为该子数组是一个 美丽 子集。

返回数组 nums 中 非空 且 美丽 的子集数目。

nums 的子集定义为：可以经由 nums 删除某些元素（也可能不删除）得到的一个数组。只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的子集。

示例 1：

输入：nums = [2,4,6], k = 2
输出：4
解释：数组 nums 中的美丽子集有：[2], [4], [6], [2, 6] 。
可以证明数组 [2,4,6] 中只存在 4 个美丽子集。
示例 2：

输入：nums = [1], k = 1
输出：1
解释：数组 nums 中的美丽数组有：[1] 。
可以证明数组 [1] 中只存在 1 个美丽子集。
*/
func beautifulSubsets(nums []int, k int) int {
	groups := map[int]map[int]int{}
	for _, x := range nums {
		if groups[x%k] == nil {
			groups[x%k] = map[int]int{}
		}
		groups[x%k][x]++
	}
	ans := 1
	for _, cnt := range groups {
		m := len(cnt)
		type pair struct{ x, c int }
		g := make([]pair, 0, m)
		for x, c := range cnt {
			g = append(g, pair{x, c})
		}
		sort.Slice(g, func(i, j int) bool { return g[i].x < g[j].x })
		f := make([]int, m+1)
		f[0] = 1
		f[1] = 1 << g[0].c
		for i := 1; i < m; i++ {
			if g[i].x-g[i-1].x == k {
				f[i+1] = f[i] + f[i-1]*(1<<g[i].c-1)
			} else {
				f[i+1] = f[i] << g[i].c
			}
		}
		ans *= f[m]
	}
	return ans - 1 // -1 去掉空集
}

func main() {
	println(beautifulSubsets([]int{2, 4, 6, 8}, 2))
}
