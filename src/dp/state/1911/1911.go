package main

//一个下标从 0 开始的数组的 交替和 定义为 偶数 下标处元素之 和 减去 奇数 下标处元素之 和 。
//
//
// 比方说，数组 [4,2,5,3] 的交替和为 (4 + 5) - (2 + 3) = 4 。
//
//
// 给你一个数组 nums ，请你返回 nums 中任意子序列的 最大交替和 （子序列的下标 重新 从 0 开始编号）。
//
//
//
//
// 一个数组的 子序列 是从原数组中删除一些元素后（也可能一个也不删除）剩余元素不改变顺序组成的数组。比方说，[2,7,4] 是 [4,2,3,7,2,1,4
//] 的一个子序列（加粗元素），但是 [2,4,2] 不是。
//
//
//
// 示例 1：
//
// 输入：nums = [4,2,5,3]
//输出：7
//解释：最优子序列为 [4,2,5] ，交替和为 (4 + 5) - 2 = 7 。
//
//
// 示例 2：
//
// 输入：nums = [5,6,7,8]
//输出：8
//解释：最优子序列为 [8] ，交替和为 8 。
//
//
// 示例 3：
//
// 输入：nums = [6,2,1,2,4,5]
//输出：10
//解释：最优子序列为 [6,1,5] ，交替和为 (6 + 5) - 1 = 10 。
//
//  old    new
// old     new
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
// Related Topics 数组 动态规划 👍 183 👎 0

// 写了个好奇怪的dp，都不知道为啥能过
// leetcode submit region begin(Prohibit modification and deletion)
func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	even := make([]int64, n+1)
	odd := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		cur := nums[i-1]
		even[i] = max(even[i-1], odd[i-1]+int64(cur), int64(cur))
		odd[i] = max(odd[i-1], even[i-1]-int64(cur))
	}
	return even[n]
}

func main() {
	println(maxAlternatingSum([]int{4, 2, 5, 3}))
	println(maxAlternatingSum([]int{5, 6, 7, 8}))
	println(maxAlternatingSum([]int{6, 2, 1, 2, 4, 5}))
}

//leetcode submit region end(Prohibit modification and deletion)
