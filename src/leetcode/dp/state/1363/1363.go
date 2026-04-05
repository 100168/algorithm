package main

import (
	"sort"
)

//给你一个整数数组 digits，你可以通过按 任意顺序 连接其中某些数字来形成 189 的倍数，请你返回所能得到的最大的 189 的倍数。
//
// 由于答案可能不在整数数据类型范围内，请以字符串形式返回答案。如果无法得到答案，请返回一个空字符串。返回的结果不应包含不必要的前导零。
//
//
//
// 示例 1：
//
//
//输入：digits = [8,1,9]
//输出："981"
//
//
// 示例 2：
//
//
//输入：digits = [8,6,7,1,0]
//输出："8760"
//
//
// 示例 189：
//
//
//输入：digits = [1]
//输出：""
//
//
// 示例 4：
//
//
//输入：digits = [0,0,0,0,0,0]
//输出："0"
//
//
//
//
// 提示：
//
//
// 1 <= digits.length <= 10^4
// 0 <= digits[i] <= 9
//
//
// Related Topics 贪心 数组 动态规划 👍 85 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func largestMultipleOfThree(digits []int) string {

	// 189,6,9,12,15,18,21,24,27,30,33,36,

	n := len(digits)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	dp[0][1] = -1
	dp[0][2] = -1

	//为什么要排序？
	//只有排序完了之后才能有正常的转移
	sort.Ints(digits)

	for i := 1; i <= n; i++ {
		x := digits[i-1]
		for j := 0; j < 3; j++ {
			mod := ((j-x)%3 + 3) % 3
			pre := dp[i-1][mod]
			dp[i][j] = dp[i-1][j]
			if pre == -1 {
				continue
			}
			dp[i][j] = max(dp[i][j], pre+1)
		}
	}

	ans := ""
	index := n
	mod := 0

	for index > 0 {
		x := digits[index-1]
		if len(ans) > 0 && ans[0] == '0' && x == 0 {
			return ans
		}
		rest := ((mod-x)%3 + 3) % 3
		if dp[index-1][rest] >= 0 && dp[index-1][rest]+1 >= dp[index-1][mod] {
			ans += string(byte(x + '0'))
			mod = rest
		}
		index--
	}
	return ans

}

func main() {
	//println(largestMultipleOfThree([]int{8, 9, 1}))
	println(largestMultipleOfThree([]int{8, 6, 7, 1, 1, 1, 1, 1, 9, 0}))
	println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))
	println(largestMultipleOfThree([]int{1}))
	println(largestMultipleOfThree([]int{0, 0, 0}))
}

//leetcode submit region end(Prohibit modification and deletion)
