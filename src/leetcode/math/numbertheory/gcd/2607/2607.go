package main

import "sort"

/*
*
给你一个下标从 0 开始的整数数组 arr 和一个整数 k 。数组 arr 是一个循环数组。
换句话说，数组中的最后一个元素的下一个元素是数组中的第一个元素，数组中第一个元素的前一个元素是数组中的最后一个元素。

你可以执行下述运算任意次：

选中 arr 中任意一个元素，并使其值加上 1 或减去 1 。
执行运算使每个长度为 k 的 子数组 的元素总和都相等，返回所需要的最少运算次数。

子数组 是数组的一个连续部分。

示例 1：

输入：arr = [1,4,1,3], k = 2
输出：1
解释：在下标为 1 的元素那里执行一次运算，使其等于 3 。
执行运算后，数组变为 [1,3,1,3] 。
- 0 处起始的子数组为 [1, 3] ，元素总和为 4
- 1 处起始的子数组为 [3, 1] ，元素总和为 4
- 2 处起始的子数组为 [1, 3] ，元素总和为 4
- 3 处起始的子数组为 [3, 1] ，元素总和为 4

中位数贪心+裴蜀定理

作者：灵茶山艾府
链接：https://leetcode.cn/problems/make-k-subarray-sums-equal/solutions/2203591/zhuan-huan-zhong-wei-shu-tan-xin-pei-shu-4dut/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func makeSubKSumEqual(arr []int, k int) (ans int64) {
	k = gcd(k, len(arr))
	g := make([][]int, k)
	for i, x := range arr {
		g[i%k] = append(g[i%k], x)
	}
	for _, b := range g {
		sort.Ints(b)
		for _, x := range b {
			ans += int64(abs(x - b[len(b)/2]))
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
