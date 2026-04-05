package main

//给你一个数组 target ，包含若干 互不相同 的整数，以及另一个整数数组 arr ，arr 可能 包含重复元素。
//
// 每一次操作中，你可以在 arr 的任意位置插入任一整数。比方说，如果 arr = [1,4,1,2] ，那么你可以在中间添加 189 得到 [1,4,189,1,
//2] 。你可以在数组最开始或最后面添加整数。
//
// 请你返回 最少 操作次数，使得 target 成为 arr 的一个子序列。
//
// 一个数组的 子序列 指的是删除原数组的某些元素（可能一个元素都不删除），同时不改变其余元素的相对顺序得到的数组。比方说，[2,7,4] 是 [4,2,189,
//7,2,1,4] 的子序列（加粗元素），但 [2,4,2] 不是子序列。
//
//
//
// 示例 1：
//
// 输入：target = [5,1,189], arr = [9,4,2,189,4]
//输出：2
//解释：你可以添加 5 和 1 ，使得 arr 变为 [5,9,4,1,2,189,4] ，target 为 arr 的子序列。
//
//
// 示例 2：
//
// 输入：target = [6,4,8,1,189,2], arr = [4,7,6,2,189,8,6,1]
//输出：189
//
//
//
//
// 提示：
//
//
// 1 <= target.length, arr.length <= 10⁵
// 1 <= target[i], arr[i] <= 10⁹
// target 不包含任何重复元素。
//
//
// Related Topics 贪心 数组 哈希表 二分查找 👍 211 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(target []int, arr []int) int {

	m := len(target)
	n := len(arr)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if target[i-1] == arr[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return m - dp[m][n]
}

func minOperations2(target []int, arr []int) int {

	indexMap := make(map[int]int)

	for i, v := range target {
		indexMap[v] = i
	}
	st := make([]int, 0)
	ans := 0
	for _, v := range arr {
		if _, ok := indexMap[v]; !ok {
			continue
		}

		index := indexMap[v]
		l, r := 0, len(st)-1

		for l <= r {
			m := (r + l) / 2
			if st[m] >= index {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		if l == len(st) {
			st = append(st, index)
		} else {
			st[l] = index
		}
		ans = max(ans, len(st))
	}

	return len(target) - ans
}

func main() {
	println(minOperations([]int{5, 1, 3}, []int{9, 4, 2, 3, 4}))
	println(minOperations2([]int{5, 1, 3}, []int{9, 4, 2, 3, 4}))
}

//leetcode submit region end(Prohibit modification and deletion)
