package main

import (
	"strings"
)

/*
*

给你两个正整数 low 和 high ，都用字符串表示，请你统计闭区间 [low, high] 内的 步进数字 数目。

如果一个整数相邻数位之间差的绝对值都 恰好 是 1 ，那么这个数字被称为 步进数字 。

请你返回一个整数，表示闭区间 [low, high] 之间步进数字的数目。

由于答案可能很大，请你将它对 109 + 7 取余 后返回。

注意：步进数字不能有前导 0 。
*/
func countSteppingNumbers(low string, high string) int {

	mod := int(1e9 + 7)

	lh := len(high)
	ll := len(low)

	low = strings.Repeat("0", lh-ll) + low

	memo := make([][]int, lh)
	for i := range memo {
		memo[i] = make([]int, 10)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool, bool) int

	dfs = func(i, pre int, isLimitL, isLimitH, isNum bool) int {
		if i == lh {
			if isNum {
				return 1
			}
			return 0
		}
		if isNum && !isLimitL && !isLimitH && memo[i][pre] != -1 {
			return memo[i][pre]
		}

		res := 0
		if !isNum && i < lh-ll {
			res += dfs(i+1, pre, true, false, false)
		}

		start := 0
		up := 9

		if isLimitL {
			start = int(low[i] - '0')
		}
		if isLimitH {
			up = int(high[i] - '0')
		}

		down := 0
		if !isNum {
			down = 1
		}

		for d := max(start, down); d <= up; d++ {
			if pre == -1 || abs(d-pre) == 1 {
				res += dfs(i+1, d, isLimitL && d == start, isLimitH && d == up, true)
			}
		}
		res %= mod
		if isNum && !isLimitH && !isLimitL {
			memo[i][pre] = res
		}

		return res
	}

	return dfs(0, -1, true, true, false)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	println(countSteppingNumbers("90", "101"))
}
