package main

/*
3098. 求出所有子序列的能量和
困难
相关企业

提示
给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。

一个整数数组的 能量 定义为和 等于 k 的子序列的数目。

请你返回 nums 中所有子序列的 能量和 。

由于答案可能很大，请你将它对 109 + 7 取余 后返回。


观察

如果和为 kkk 的子序列 SSS 的长度是 ccc，那么有多少个子序列 TTT，会包含 SSS 呢？即 SSS 是 TTT 的子序列。

例如示例 1，子序列 S=[3]S=[3]S=[3]，出现在子序列 [1,2,3],[1,3],[2,3],[3][1,2,3],[1,3],[2,3],[3][1,2,3],[1,3],[2,3],[3] 中，这 444 个子序列都可以是 TTT。

除了 333 以外的每个数，都可以选择在/不在包含 [3][3][3] 的子序列 TTT 中。

所以有 2n−c2^{n-c}2
n−c
  个子序列 TTT。

这意味着 SSS 对答案的贡献是 2n−c2^{n-c}2
n−c
 。

作者：灵茶山艾府
链接：https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences/solutions/2691661/liang-chong-fang-fa-er-wei-yi-wei-0-1-be-2e47/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

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
// Related Topics 数组 动态规划 👍 7 👎 0

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
	dfs = func(index, rest int) int {
		if rest < 0 {
			return 0
		}

		if index == n {
			if rest == 0 {
				return 1
			}
			return 0
		}
		if cache[index][rest] != -1 {
			return cache[index][rest]
		}

		cnt := 0
		cnt += dfs(index+1, rest)*2%mod + dfs(index+1, rest-nums[index])%mod
		cache[index][rest] = cnt
		return cnt % mod
	}
	return dfs(0, k)
}

func main() {
	println(sumOfPower([]int{3, 4}, 2))
}
