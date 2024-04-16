package main

import "slices"

//给定由
// n 个小写字母字符串组成的数组
// strs ，其中每个字符串长度相等。
//
// 选取一个删除索引序列，对于
// strs 中的每个字符串，删除对应每个索引处的字符。
//
// 比如，有
// strs = ["abcdef","uvwxyz"] ，删除索引序列
// {0, 2, 3} ，删除后为
// ["bef", "vyz"] 。
//
// 假设，我们选择了一组删除索引
// answer ，那么在执行删除操作之后，最终得到的数组的行中的 每个元素 都是按字典序排列的（即 (strs[0][0] <= strs[0][1] <=
// ... <= strs[0][strs[0].length - 1]) 和 (strs[1][0] <= strs[1][1] <= ... <= strs[
//1][strs[1].length - 1]) ，依此类推）。
//
// 请返回
// answer.length 的最小可能值 。
//
//
//
// 示例 1：
//
//
//输入：strs = ["babca","bbazb"]
//输出：3
//解释：
//删除 0、1 和 4 这三列后，最终得到的数组是 strs = ["bc", "az"]。
//这两行是分别按字典序排列的（即，strs[0][0] <= strs[0][1] 且 strs[1][0] <= strs[1][1]）。
//注意，strs[0] > strs[1] —— 数组 strs 不一定是按字典序排列的。
//
//
// 示例 2：
//
//
//输入：strs = ["edcba"]
//输出：4
//解释：如果删除的列少于 4 列，则剩下的行都不会按字典序排列。
//
//
// 示例 3：
//
//
//输入：strs = ["ghi","def","abc"]
//输出：0
//解释：所有行都已按字典序排列。
//
//
//
//
// 提示：
//
//
//
// n == strs.length
// 1 <= n <= 100
// 1 <= strs[i].length <= 100
// strs[i] 由小写英文字母组成
//
//
// Related Topics 数组 字符串 动态规划 👍 78 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minDeletionSize(strs []string) int {
	m := len(strs)
	n := len(strs[0])
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			flag := true
			for k := 0; k < m; k++ {
				if strs[k][j-1] > strs[k][i-1] {
					flag = false
				}
			}
			if flag {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return n - slices.Max(dp)
}

//leetcode submit region end(Prohibit modification and deletion)
