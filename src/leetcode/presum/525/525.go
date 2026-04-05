package main

/*
525. 连续数组
中等
相关标签
相关企业
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2:

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。

s[r]-s[l] = 0
*/
func findMaxLength(nums []int) int {
	cnt := make(map[int]int)
	cnt[0] = -1

	s := 0
	ans := 0
	for i, num := range nums {
		s += (num * 2) - 1

		if v, ok := cnt[s]; ok {
			ans = max(ans, i-v)
		}

		if _, ok := cnt[s]; !ok {
			cnt[s] = i
		}

	}
	return ans
}
