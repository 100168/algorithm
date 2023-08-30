package main

import (
	"strconv"
)

//给定一个按 非递减顺序 排列的数字数组
// digits 。你可以用任意次数 digits[i] 来写的数字。例如，如果
// digits = ['1','3','5']，我们可以写数字，如
// '13', '551', 和 '1351315'。
//
// 返回 可以生成的小于或等于给定整数 n 的正整数的个数 。
//

func atMostNGivenDigitSet(digits []string, n int) int {

	s := strconv.Itoa(n)
	n = len(s)
	memo := make([]int, n)

	for i := range memo {
		memo[i] = -1
	}
	return dfs(0, true, false, memo, s, digits)

}

func dfs(i int, isLimit bool, isNum bool, memo []int, s string, digits []string) int {

	if i == len(s) {
		if isNum {
			return 1
		}
		return 0
	}
	if !isLimit && isNum && memo[i] != -1 {
		return memo[i]
	}
	res := 0
	if !isNum {
		res = dfs(i+1, false, false, memo, s, digits)
	}

	up := 9
	if isLimit {
		up = int(s[i] - '0')
	}
	for _, v := range digits {
		cur := int(v[0] - '0')
		if cur <= up {
			res += dfs(i+1, isLimit && cur == up, true, memo, s, digits)
		}
	}
	if !isLimit && isNum {
		memo[i] = res
	}
	return res
}
