package _33

import "strconv"

// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
// 100
// 1,11,12,13,14,15,16,17,
func countDigitOne(n int) int {

	s := strconv.Itoa(n)
	l := len(s)
	memo := make([][]int, l)
	for i := range memo {
		memo[i] = make([]int, (1<<l)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dfs(0, 0, true, false, s, memo)
}

func dfs(i int, pre int, isLimit bool, isNum bool, s string, memo [][]int) int {

	if i == len(memo) {

		return pre
	}

	if !isLimit && isNum && memo[i][pre] != -1 {
		return memo[i][pre]
	}

	res := 0

	if !isNum {
		res = dfs(i+1, pre, false, false, s, memo)
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

		if j == 1 {
			pre++
		}
		res += dfs(i+1, pre, isLimit && j == up, true, s, memo)
		if j == 1 {
			pre--
		}
	}
	if !isLimit && isNum {
		memo[i][pre] = res
	}
	return res
}
