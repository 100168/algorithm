package main

import (
	"strconv"
)

//给定一个正整数 n ，请你统计在 [0, n] 范围的非负整数中，有多少个整数的二进制表示中不存在 连续的 1 。
/*
9  -> 1001
10 -> 1010
11 -> 1011
12 -> 1100

*/

func findIntegers(n int) int {

	s := strconv.FormatInt(int64(n), 2)

	m := len(s)
	memo := make([][]int, m)

	for i := range memo {
		memo[i] = make([]int, 2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int

	dfs = func(i, pre int, isLimit bool, isNum bool) int {
		if i == m {
			return 1
		}
		if !isLimit && isNum && memo[i][pre] != -1 {
			return memo[i][pre]
		}
		res := 0
		if !isNum {
			res += dfs(i+1, 0, false, false)
		}
		up := 1
		if isLimit {
			up = int(s[i] - '0')
		}
		low := 0
		if !isNum {
			low = 1
		}
		for j := low; j <= up; j++ {
			if pre != 1 || j != 1 {
				res += dfs(i+1, j, isLimit && j == up, true)
			}
		}
		if !isLimit && isNum {
			memo[i][pre] = res
		}
		return res
	}
	return dfs(0, 0, true, false)
}
func main() {
	println(findIntegers(5))
}
