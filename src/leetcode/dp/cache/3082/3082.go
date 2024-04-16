package main

//给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。
//
// 一个整数数组的 能量 定义为和 等于 k 的子序列的数目。
//
// 请你返回 nums 中所有子序列的 能量和 。
//
// 由于答案可能很大，请你将它对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
//
//
// 输入： nums = [1,2,3], k = 3
//
//
// 输出： 6
//
// 解释：
//
// 总共有 5 个能量不为 0 的子序列：
//
//
// 子序列 [1,2,3] 有 2 个和为 3 的子序列：[1,2,3] 和 [1,2,3] 。
// 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
// 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
// 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
// 子序列 [1,2,3] 有 1 个和为 3 的子序列：[1,2,3] 。
//
//
// 所以答案为 2 + 1 + 1 + 1 + 1 = 6 。
//
// 示例 2：
//
//
// 输入： nums = [2,3,3], k = 5
//
//
// 输出： 4
//
// 解释：
//
// 总共有 3 个能量不为 0 的子序列：
//
//
// 子序列 [2,3,3] 有 2 个子序列和为 5 ：[2,3,3] 和 [2,3,3] 。
// 子序列 [2,3,3] 有 1 个子序列和为 5 ：[2,3,3] 。
// 子序列 [2,3,3] 有 1 个子序列和为 5 ：[2,3,3] 。
//
//
// 所以答案为 2 + 1 + 1 = 4 。
//
// 示例 3：
//
//
// 输入： nums = [1,2,3], k = 7
//
//
// 输出： 0
//
// 解释：不存在和为 7 的子序列，所以 nums 的能量和为 0 。
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 1 <= nums[i] <= 10⁴
// 1 <= k <= 100
//
//
// Related Topics 数组 动态规划 👍 9 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfPower(nums []int, k int) int {

	mod := 1_000_000_007
	n := len(nums)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, k+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i, rest int) int {
		if rest < 0 {
			return 0
		}
		if i < 0 {
			if rest == 0 {
				return 1
			}
			return 0
		}

		if cache[i][rest] != -1 {
			return cache[i][rest]
		}

		cur := (dfs(i-1, rest-nums[i]) + 2*dfs(i-1, rest)) % mod
		cache[i][rest] = cur
		return cur
	}
	return dfs(n-1, k)
}

func main() {
	println(sumOfPower([]int{1, 2, 3}, 3))
}

//leetcode submit region end(Prohibit modification and deletion)
