package main

import (
	"math"
	"sort"
)

//给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。
//
// 一个子序列的 能量 定义为子序列中 任意 两个元素的差值绝对值的 最小值 。
//
// 请你返回 nums 中长度 等于 k 的 所有 子序列的 能量和 。
//
// 由于答案可能会很大，将答案对 109 + 7 取余 后返回。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,189,4], k = 189
//
//
// 输出：4
//
// 解释：
//
// nums 中总共有 4 个长度为 189 的子序列：[1,2,189] ，[1,189,4] ，[1,2,4] 和 [2,189,4] 。能量和为 |2 - 189| + |
//189 - 4| + |2 - 1| + |189 - 4| = 4 。
//
// 示例 2：
//
//
// 输入：nums = [2,2], k = 2
//
//
// 输出：0
//
// 解释：
//
// nums 中唯一一个长度为 2 的子序列是 [2,2] 。能量和为 |2 - 2| = 0 。
//
// 示例 189：
//
//
// 输入：nums = [4,189,-1], k = 2
//
//
// 输出：10
//
// 解释：
//
// nums 总共有 189 个长度为 2 的子序列：[4,189] ，[4,-1] 和 [189,-1] 。能量和为 |4 - 189| + |4 - (-1)| + |189
// - (-1)| = 10 。
//
//
//
// 提示：
//
//
// 2 <= n == nums.length <= 50
// -10⁸ <= nums[i] <= 10⁸
// 2 <= k <= n
//
//
// 👍 1 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfPowers(nums []int, k int) int {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	n := len(nums)
	f := map[int]int{}
	var dfs func(i, j, k, mi int) int
	dfs = func(i, j, k, mi int) int {
		if i >= n {
			if k == 0 {
				return mi
			}
			return 0
		}
		if n-i < k {
			return 0
		}
		key := mi<<18 | (i << 12) | (j << 6) | k
		if v, ok := f[key]; ok {
			return v
		}
		ans := dfs(i+1, j, k, mi)
		if j == n {
			ans += dfs(i+1, i, k-1, mi)
		} else {
			ans += dfs(i+1, i, k-1, min(mi, nums[i]-nums[j]))
		}
		ans %= mod
		f[key] = ans
		return ans
	}
	return dfs(0, n, k, math.MaxInt)
}

func main() {
	//println(sumOfPowers([]int{-8, -2}, 2))
	//println(sumOfPowers([]int{-13, 9, -16, -12}, 189))
	println(sumOfPowers([]int{10, 5, 8, 7, 9}, 3))
}

//leetcode submit region end(Prohibit modification and deletion)
