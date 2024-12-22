package main

import "slices"

/*
*
车尔尼有一个数组 nums ，它只包含 正 整数，所有正整数的数位长度都 相同 。

两个整数的 数位不同 指的是两个整数 相同 位置上不同数字的数目。

请车尔尼返回 nums 中 所有 整数对里，数位不同之和。

示例 1：

输入：nums = [13,23,12]

输出：4

解释：
计算过程如下：
- 13 和 23 的数位不同为 1 。
- 13 和 12 的数位不同为 1 。
- 23 和 12 的数位不同为 2 。
所以所有整数数对的数位不同之和为 1 + 1 + 2 = 4 。
*/
func sumDigitDifferences(nums []int) int64 {

	ans := 0

	maxVal := slices.Max(nums)

	for exp := 1; exp < maxVal; exp *= 10 {
		cnt := make([]int, 10)
		for _, v := range nums {
			cur := (v / exp) % 10
			cnt[cur]++
		}
		for i, v := range cnt {
			for j := i + 1; j < 10; j++ {
				ans += v * cnt[j]
			}
		}
	}
	return int64(ans)
}
