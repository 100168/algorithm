package main

import (
	"fmt"
	"slices"
)

/*
*
给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：

有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
返回 图中最大连通组件的大小 。
*/
func largestComponentSize(nums []int) int {

	n := len(nums)
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	find := func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) {
		fa := find(a)
		fb := find(b)
		parent[fa] = fb
	}

	maxVal := slices.Max(nums)
	indexNums := make([]int, maxVal+1)

	for i := range indexNums {
		indexNums[i] = -1
	}
	for j, v := range nums {

		for i := 2; i*i <= v; i++ {
			if v%i != 0 {
				continue
			}
			for v%i == 0 {
				v /= i
			}
			if indexNums[i] != -1 {
				union(indexNums[i], j)
			}
			indexNums[i] = j
		}
		if v > 1 {
			if indexNums[v] != -1 {
				union(indexNums[v], j)
			}
			indexNums[v] = j
		}

	}

	ans := 0
	for i := range nums {
		size[find(i)]++
		ans = max(ans, size[parent[i]])
	}
	return ans
}

func main() {
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
}
