package main

import (
	"slices"
	"sort"
)

//给你一个整数数组 digits，你可以通过按 任意顺序 连接其中某些数字来形成 3 的倍数，请你返回所能得到的最大的 3 的倍数。
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
// 示例 3：
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

	// 3,6,9,12,15,18,21,24,27,30,33,36,

	n := len(digits)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	dp[0][1] = -1
	dp[0][2] = -1

	sort.Ints(digits)

	path := make([][][]byte, 2)
	for i := range path {
		path[i] = make([][]byte, 3)
		for j := range path[i] {
			path[i][j] = make([]byte, 0)
		}
	}

	for i, x := range digits {
		for j := 0; j < 3; j++ {
			pre := dp[i%2][((j-x)%3+3)%3]
			dp[(i+1)%2][j] = dp[i%2][j]
			nP := make([]byte, len(path[i%2][j]))
			copy(nP, path[i%2][j])
			path[(i+1)%2][j] = nP
			if pre == -1 {
				continue
			}
			dp[(i+1)%2][j] = max(dp[(i)%2][j], pre+1)
			if dp[(i+1)%2][j] == pre+1 {
				nP := make([]byte, len(path[i%2][((j-x)%3+3)%3]))
				copy(nP, path[i%2][((j-x)%3+3)%3])
				nP = append(nP, byte(x+'0'))
				path[(i+1)%2][j] = nP
			}
		}
	}
	slices.Sort(path[n%2][0])
	slices.Reverse(path[n%2][0])

	if len(path[n%2][0]) > 0 && path[n%2][0][0] == '0' {

		return string('0')
	}
	return string(path[n%2][0])

}

func main() {
	println(largestMultipleOfThree([]int{8, 9, 1}))
	println(largestMultipleOfThree([]int{8, 6, 7, 1, 1, 1, 1, 1, 9, 0}))
	println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))
}

//leetcode submit region end(Prohibit modification and deletion)
