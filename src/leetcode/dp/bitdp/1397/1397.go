package main

import "strings"

var mod = 1_000_000_007

/*
*
给你两个长度为 n 的字符串 s1 和 s2 ，以及一个字符串 evil 。请你返回 好字符串 的数目。

好字符串 的定义为：它的长度为 n ，字典序大于等于 s1 ，字典序小于等于 s2 ，且不包含 evil 为子字符串。

由于答案可能很大，请你返回答案对 10^9 + 7 取余的结果。
*/
func findGoodStrings(n int, s1 string, s2 string, evil string) int {

	next := getNext(evil)
	flag := 0
	if strings.Index(s1, evil) == -1 {
		flag = 1
	}

	endV := calc(s2, n, next, evil)
	startV := calc(s1, n, next, evil)
	v := endV - startV
	if v < 0 {
		v += mod
	}

	return v + flag

}

func getNext(evil string) []int {

	n := len(evil)
	//下一个匹配的位置
	next := make([]int, n)
	for i := 1; i < n; i++ {
		j := next[i-1]
		for j > 0 && evil[i] != evil[j] {
			j = next[j-1]
		}
		if evil[i] == evil[j] {
			next[i] = j + 1
		}
	}
	return next
}

func calc(end string, n int, next []int, evil string) int {

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, len(next))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dfs(0, 0, true, next, memo, end, evil)

}

func dfs(i int, pre int, isLimit bool, next []int, memo [][]int, s string, evil string) int {

	if pre == len(next) {
		return 0
	}
	if i == len(memo) {
		return 1
	}
	if !isLimit && memo[i][pre] != -1 {
		return memo[i][pre]
	}

	up := uint8('z')
	if isLimit {
		up = s[i]
	}
	res := 0
	for end := uint8('a'); end <= up; end++ {
		curPre := pre
		for curPre != 0 && evil[curPre] != end {
			curPre = next[curPre-1]
		}
		if evil[curPre] == end {
			curPre++
		}
		res += dfs(i+1, curPre, isLimit && end == up, next, memo, s, evil) % mod
		res %= mod
	}
	if !isLimit {
		memo[i][pre] = res % mod
	}
	return res
}

func main() {
	println(findGoodStrings(2, "aa", "da", "b"))
}
