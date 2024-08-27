package main

import (
	"math"
	"sort"
)

/**
给你一个由 n 个数对组成的数对数组 pairs ，其中 pairs[i] = [lefti, righti] 且 lefti < righti 。

现在，我们定义一种 跟随 关系，当且仅当 b < c 时，数对 p2 = [c, d] 才可以跟在 p1 = [a, b] 后面。我们用这种形式来构造 数对链 。

找出并返回能够形成的 最长数对链的长度 。

你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。



示例 1：

输入：pairs = [[1,2], [2,3], [3,4]]
输出：2
解释：最长的数对链是 [1,2] -> [3,4] 。
示例 2：

输入：pairs = [[1,2],[7,8],[4,5]]
输出：3
解释：最长的数对链是 [1,2] -> [4,5] -> [7,8] 。

*/

func findLongestChain(pairs [][]int) int {

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	end := math.MinInt

	ans := 0
	for _, v := range pairs {
		if v[0] > end {
			end = v[1]
			ans++
		}
	}
	return ans

}
