package main

import "strconv"

// 给定一个在 0 到 9 之间的整数 d，和两个正整数 low 和 high 分别作为上下界。返回 d 在 low 和 high 之间的整数中出现的次数，包括
// 边界 low 和 high。
// 示例 1：
// 输入：d = 1, low = 1, high = 13
// 输出：6
// 解释：
// 数字 d=1 在 1,10,11,12,13 中出现 6 次。注意 d=1 在数字 11 中出现两次。
func digitsCount(d int, low int, high int) int {

	return countDigitOne(high, d) - countDigitOne(low-1, d)
}

func countDigitOne(n int, d int) int {

	s := strconv.Itoa(n)
	l := len(s)
	memo := make([][]int, l)
	for i := range memo {
		memo[i] = make([]int, (1<<l)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dfs(0, 0, true, false, s, memo, d)
}

func dfs(i int, pre int, isLimit bool, isNum bool, s string, memo [][]int, d int) int {

	if i == len(memo) {

		return pre
	}

	if !isLimit && isNum && memo[i][pre] != -1 {
		return memo[i][pre]
	}

	res := 0

	if !isNum {
		res = dfs(i+1, pre, false, false, s, memo, d)
	}
	up := 9
	start := 0
	if isLimit {
		up = int(s[i] - '0')
	}
	if !isNum {
		start = 1
	}

	for j := start; j <= up; j++ {

		if j == d {
			pre++
		}
		res += dfs(i+1, pre, isLimit && j == up, true, s, memo, d)
		if j == d {
			pre--
		}
	}
	if !isLimit && isNum {
		memo[i][pre] = res
	}
	return res
}
