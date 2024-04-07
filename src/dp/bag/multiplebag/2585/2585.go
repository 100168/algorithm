package main

import "fmt"

//考试中有 n 种类型的题目。给你一个整数 target 和一个下标从 0 开始的二维整数数组 types ，其中 types[i] = [counti,
//marksi] 表示第 i 种类型的题目有 counti 道，每道题目对应 marksi 分。
//
// 返回你在考试中恰好得到 target 分的方法数。由于答案可能很大，结果需要对 10⁹ +7 取余。
//
// 注意，同类型题目无法区分。
//
//
// 比如说，如果有 3 道同类型题目，那么解答第 1 和第 2 道题目与解答第 1 和第 3 道题目或者第 2 和第 3 道题目是相同的。
//
//
//
//
// 示例 1：
//
// 输入：target = 6, types = [[6,1],[3,2],[2,3]]
//输出：7
//解释：要获得 6 分，你可以选择以下七种方法之一：
//- 解决 6 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 + 1 = 6
//- 解决 4 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 1 + 2 = 6
//- 解决 2 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 1 + 2 + 2 = 6
//- 解决 3 道第 0 种类型的题目和 1 道第 2 种类型的题目：1 + 1 + 1 + 3 = 6
//- 解决 1 道第 0 种类型的题目、1 道第 1 种类型的题目和 1 道第 2 种类型的题目：1 + 2 + 3 = 6
//- 解决 3 道第 1 种类型的题目：2 + 2 + 2 = 6
//- 解决 2 道第 2 种类型的题目：3 + 3 = 6
//
//
// 示例 2：
//
// 输入：target = 5, types = [[50,1],[50,2],[50,5]]
//输出：4
//解释：要获得 5 分，你可以选择以下四种方法之一：
//- 解决 5 道第 0 种类型的题目：1 + 1 + 1 + 1 + 1 = 5
//- 解决 3 道第 0 种类型的题目和 1 道第 1 种类型的题目：1 + 1 + 1 + 2 = 5
//- 解决 1 道第 0 种类型的题目和 2 道第 1 种类型的题目：1 + 2 + 2 = 5
//- 解决 1 道第 2 种类型的题目：5
//
//
// 示例 3：
//
// 输入：target = 18, types = [[6,1],[3,2],[2,3]]
//输出：1
//解释：只有回答所有题目才能获得 18 分。
//
//
//
//
// 提示：
//
//
// 1 <= target <= 1000
// n == types.length
// 1 <= n <= 50
// types[i].length == 2
// 1 <= counti, marksi <= 50
//
//
// Related Topics 数组 动态规划 👍 40 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func waysToReachTarget(target int, types [][]int) int {

	mod := 1_000_000_007

	n := len(types)
	var dfs func(int, int) int
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dfs = func(i, rest int) int {
		if i < 0 {
			if rest == 0 {
				return 1
			}
			return 0
		}
		if dp[i][rest] != -1 {
			return dp[i][rest]
		}
		c, m := types[i][0], types[i][1]
		cur := 0
		for k := 0; k <= c && m*k <= rest; k++ {
			cur += dfs(i-1, rest-m*k)
			cur %= mod
		}
		dp[i][rest] = cur
		return cur
	}
	return dfs(n-1, target)
}

func dp(target int, types [][]int) int {

	mod := 1_000_000_007

	mark := make([]int, 0)
	for _, v := range types {
		c, m := v[0], v[1]
		i := 0
		for 1<<i < c {
			mark = append(mark, 1<<i*m)
			c -= 1 << i
			i++
		}
		if c > 0 {
			mark = append(mark, c*m)
		}
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for _, v := range mark {
		for j := target; j >= v; j-- {
			dp[j] += dp[j-v]
			dp[j] %= mod
		}
	}
	return dp[target]
}

func waysToReachTarget2(target int, types [][]int) int {
	const mod int = 1e9 + 7
	f := make([]int, target+1)
	f[0] = 1
	for _, p := range types {
		count, marks := p[0], p[1]
		for j := target; j > 0; j-- {
			for k := 1; k <= count && k*marks <= j; k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}
func main() {
	fmt.Println(waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))
	fmt.Println(dp(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))
	fmt.Println(waysToReachTarget2(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))
}

//leetcode submit region end(Prohibit modification and deletion)
