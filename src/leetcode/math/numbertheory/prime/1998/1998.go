package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
*
给你一个整数数组 nums ，你可以在 nums 上执行下述操作 任意次 ：

如果 gcd(nums[i], nums[j]) > 1 ，交换 nums[i] 和 nums[j] 的位置。
其中 gcd(nums[i], nums[j]) 是 nums[i] 和 nums[j] 的最大公因数。
如果能使用上述交换方式将 nums 按 非递减顺序 排列，返回 true ；否则，返回 false 。

示例 1：

输入：nums = [7,21,3]
输出：true
解释：可以执行下述操作完成对 [7,21,3] 的排序：
- 交换 7 和 21 因为 gcd(7,21) = 7 。nums = [21,7,3]
- 交换 21 和 3 因为 gcd(21,3) = 3 。nums = [3,7,21]
示例 2：

输入：nums = [5,2,6,2]
输出：false
解释：无法完成排序，因为 5 不能与其他元素交换。
示例 3：

输入：nums = [10,5,9,3,15]
输出：true
解释：
可以执行下述操作完成对 [10,5,9,3,15] 的排序：
- 交换 10 和 15 因为 gcd(10,15) = 5 。nums = [15,5,9,3,10]
- 交换 15 和 3 因为 gcd(15,3) = 3 。nums = [3,5,9,15,10]
- 交换 10 和 15 因为 gcd(10,15) = 5 。nums = [3,5,9,10,15]

提示：

1 <= nums.length <= 3 * 104
2 <= nums[i] <= 105
*/
func gcdSort(nums []int) bool {

	n := len(nums)
	uf := new(uniFind)
	uf.init(n)

	numsMap := make(map[int]int)

	for i, v := range nums {

		for j := 2; j*j <= v; j++ {
			if v%j == 0 {
				if e, ok := numsMap[j]; ok {
					uf.unit(i, e)
				}
				numsMap[j] = i
			}
			for v%j == 0 {
				v /= j
			}

		}

		if v > 1 {
			if e, ok := numsMap[v]; ok {
				uf.unit(i, e)
			}
			numsMap[v] = i
		}
	}

	clone := slices.Clone(nums)
	group := uf.group
	for k, v := range group {
		cc := slices.Clone(v)
		sort.Ints(cc)
		sort.Slice(v, func(i, j int) bool {
			return nums[v[i]] < nums[v[j]]
		})
		for i, c := range group[k] {
			clone[cc[i]] = nums[c]
		}
	}

	for i, v := range clone[1:] {

		if v < clone[i] {
			return false
		}

	}
	return true

}

type uniFind struct {
	parent []int

	group map[int][]int
}

func (uniFind *uniFind) init(n int) {
	uniFind.parent = make([]int, n)
	uniFind.group = make(map[int][]int)
	for i := range uniFind.parent {
		uniFind.parent[i] = i
		uniFind.group[i] = append(uniFind.group[i], i)
	}
}

func (uniFind *uniFind) find(v int) int {
	parent := uniFind.parent
	for parent[v] != v {
		parent[v] = parent[parent[v]]
		v = parent[v]
	}
	return v
}

func (uniFind *uniFind) unit(a, b int) {
	findA := uniFind.find(a)
	findB := uniFind.find(b)
	if findA == findB {
		return
	}
	uniFind.parent[findA] = findB
	uniFind.group[findB] = append(uniFind.group[findB], uniFind.group[findA]...)
	delete(uniFind.group, findA)
}

func main() {
	//fmt.Println(gcdSort([]int{7, 21, 3}))
	//fmt.Println(gcdSort([]int{5, 2, 6, 2}))
	fmt.Println(gcdSort([]int{10, 5, 9, 3, 15}))
}
