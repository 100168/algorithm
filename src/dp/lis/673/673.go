package main

import "slices"

//给定一个未排序的整数数组
// nums ， 返回最长递增子序列的个数 。
//
// 注意 这个数列必须是 严格 递增的。
//
//
//
// 示例 1:
//
//
//输入: [1,3,5,4,7]
//输出: 2
//解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//
//
// 示例 2:
//
//
//输入: [2,2,2,2,2]
//输出: 5
//解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
//
//
//
// 提示:
//
//
//
//
//
// 1 <= nums.length <= 2000
// -10⁶ <= nums[i] <= 10⁶
//
//
// Related Topics 树状数组 线段树 数组 动态规划 👍 856 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func findNumberOfLIS(nums []int) int {

	m := len(nums)
	dp := make([]int, m+1)
	cnt := make([]int, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] == 1 {
			cnt[i]++
		}
	}

	ans := 0
	maxVal := slices.Max(dp)
	for i, v := range dp {
		if v == maxVal {
			ans += cnt[i]
		}
	}
	return ans
}

type bitSet struct {
	sum []int
	len int
}

func lowBit(index int) int {
	return index & -index
}
func (b *bitSet) query(index int) int {

	s := 0
	for index > 0 {
		s += b.sum[index]
		index -= lowBit(index)
	}
	return s
}

func (b *bitSet) update(index int) {
	for b.len > index {
		b.sum[index]++
		index += lowBit(index)
	}
}
func main() {
	println(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
}

//leetcode submit region end(Prohibit modification and deletion)
