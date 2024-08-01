package main

import (
	"fmt"
	"sort"
)

/**
给你一个下标从 0 开始的 正整数 数组 nums 和一个 正整数 limit 。

在一次操作中，你可以选择任意两个下标 i 和 j，如果 满足 |nums[i] - nums[j]| <= limit ，则交换 nums[i] 和 nums[j] 。

返回执行任意次操作后能得到的 字典序最小的数组 。

如果在数组 a 和数组 b 第一个不同的位置上，数组 a 中的对应元素比数组 b 中的对应元素的字典序更小，则认为数组 a 就比数组 b 字典序更小。例如，数组 [2,10,3] 比数组 [10,2,3] 字典序更小，下标 0 处是两个数组第一个不同的位置，且 2 < 10 。



示例 1：

输入：nums = [1,5,3,9,8], limit = 2
输出：[1,3,5,8,9]
解释：执行 2 次操作：
- 交换 nums[1] 和 nums[2] 。数组变为 [1,3,5,9,8] 。
- 交换 nums[3] 和 nums[4] 。数组变为 [1,3,5,8,9] 。
即便执行更多次操作，也无法得到字典序更小的数组。
注意，执行不同的操作也可能会得到相同的结果。

输入：nums = [1,7,28,19,10], limit = 3
输出：[1,7,28,19,10]
解释：[1,7,28,19,10] 是字典序最小的数组，因为不管怎么选择下标都无法执行操作。
*/

func lexicographicallySmallestArray(nums []int, limit int) []int {

	n := len(nums)

	type pair struct {
		index int
		val   int
	}
	pairs := make([]pair, n)

	for i, v := range nums {
		pairs[i] = pair{i, v}
	}
	groupMap := make(map[pair][]pair)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	first := pairs[0]
	groupMap[first] = append(groupMap[first], first)
	for i := 1; i < n; i++ {
		cur := pairs[i]
		pre := pairs[i-1]
		if cur.val-pre.val > limit {
			first = cur
		}
		groupMap[first] = append(groupMap[first], cur)
	}

	for _, v := range groupMap {
		indexs := make([]int, 0)
		for _, index := range v {
			indexs = append(indexs, index.index)
		}
		sort.Ints(indexs)

		for i, index := range indexs {
			nums[index] = v[i].val
		}

	}
	return nums
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(lexicographicallySmallestArray([]int{1, 7, 28, 19, 10}, 3))
	fmt.Println(lexicographicallySmallestArray([]int{1, 5, 3, 9, 8}, 3))
}
