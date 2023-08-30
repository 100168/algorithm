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
	l := len(s)
	memo := make([][]int, l)
	for i := range memo {
		memo[i] = make([]int, 2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dfs(0, false, true, false, memo, s)
}

func dfs(i int, isOne bool, isLimit bool, isNum bool, memo [][]int, s string) int {

	//base case
	if i == len(s) {
		return 1
	}
	j := 0
	if isOne {
		j = 1
	}
	//是否计算过
	if !isLimit && isNum && memo[i][j] != -1 {
		return memo[i][j]
	}

	res := 0
	//跳过当前数
	if !isNum {
		res = dfs(i+1, false, false, false, memo, s)
	}
	up := 1
	start := 0
	//受限制
	if isLimit {
		up = int(s[i] - '0')
	}
	//当前是否填过数
	if !isNum {
		start = 1
	}

	for end := start; end <= up; end++ {
		if !isOne || isOne && end != 1 {

			res += dfs(i+1, end == 1, isLimit && end == up, true, memo, s)
		}

	}
	if !isLimit && isNum {
		memo[i][j] = res
	}
	return res
}

func main() {
	println(findIntegers(5))
}
