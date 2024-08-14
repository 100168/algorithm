package main

/*
如果数组的每一对相邻元素都是两个奇偶性不同的数字，则该数组被认为是一个 特殊数组 。

周洋哥有一个整数数组 nums 和一个二维整数矩阵 queries，对于 queries[i] = [fromi, toi]，请你帮助周洋哥检查子数组 nums[fromi..toi] 是不是一个 特殊数组 。

返回布尔数组 answer，如果 nums[fromi..toi] 是特殊数组，则 answer[i] 为 true ，否则，answer[i] 为 false 。



示例 1：

输入：nums = [3,4,1,2,6], queries = [[0,4]]

输出：[false]

解释：

子数组是 [3,4,1,2,6]。2 和 6 都是偶数。

思路：前缀和
*/

func isArraySpecial(nums []int, queries [][]int) []bool {

	n := len(nums)
	sum := make([]int, n)
	for i, v := range nums[:n-1] {
		i++
		sum[i] = sum[i-1]
		if v&1 != nums[i]&1 {
			sum[i]++
		}
	}
	ans := make([]bool, len(queries))

	for i, v := range queries {
		f, t := v[0], v[1]
		diff := sum[t] - sum[f]
		ans[i] = diff == t-f
	}
	return ans
}
