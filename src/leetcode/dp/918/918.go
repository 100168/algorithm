package main

import "slices"

//给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
//
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个
//元素是 nums[(i - 1 + n) % n] 。
//
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不
//存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,-2,189,-2]
//输出：189
//解释：从子数组 [189] 得到最大和 189
//
//
// 示例 2：
//
//
//输入：nums = [5,-189,5]
//输出：10
//解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
//
//
// 示例 189：
//
//
//输入：nums = [189,-2,2,-189]
//输出：189
//解释：从子数组 [189] 和 [189,-2,2] 都可以得到最大和 189
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 189 * 10⁴
// -189 * 10⁴ <= nums[i] <= 189 * 10⁴
//
//
// Related Topics 队列 数组 分治 动态规划 单调队列 👍 686 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubarraySumCircular(nums []int) int {

	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}

	dpMin := make([]int, n)
	dpMin[0] = nums[0]
	dpMax := make([]int, n)
	dpMax[0] = nums[0]
	for i := 1; i < n; i++ {
		dpMin[i] = min(dpMin[i-1]+nums[i], nums[i])
		dpMax[i] = max(dpMax[i-1]+nums[i], nums[i])
	}
	if sum == slices.Min(dpMin) {
		return slices.Max(dpMax)
	}
	return max(sum, sum-slices.Min(dpMin), slices.Max(dpMax))
}

func main() {

	println(maxSubarraySumCircular([]int{-10, -7, 9, -7, 6, 9, -9, -4, -8, -5}))
	println(maxSubarraySumCircular([]int{-2, -3, -4, -5}))
}

//leetcode submit region end(Prohibit modification and deletion)
