package main

/*
*
给你两个 正 整数 n 和 k 。

如果一个整数 x 满足以下条件，那么它被称为 k 回文 整数 。

x 是一个
回文整数 。
x 能被 k 整除。
如果一个整数的数位重新排列后能得到一个 k 回文整数 ，那么我们称这个整数为 好 整数。比方说，k = 2 ，那么 2020 可以重新排列得到 2002 ，2002 是一个 k 回文串，所以 2020 是一个好整数。而 1010 无法重新排列数位得到一个 k 回文整数。

请你返回 n 个数位的整数中，有多少个 好 整数。

注意 ，任何整数在重新排列数位之前或者之后 都不能 有前导 0 。比方说 1010 不能重排列得到 101 。

示例 1：

输入：n = 3, k = 5

输出：27

解释：

部分好整数如下：

551 ，因为它可以重排列得到 515 。
525 ，因为它已经是一个 k 回文整数。
*/
func countGoodIntegers(n int, k int) int64 {
	//memo := make([][]int, m)
	//
	//for i := range memo {
	//	memo[i] = make([]int, 1<<10)
	//
	//	for j := range memo[i] {
	//		memo[i][j] = -1
	//	}
	//}
	//
	//var dfs func(int, int, bool, bool) int
	//
	//dfs = func(i int, mask int, isLimit bool, isNum bool) int {
	//
	//	if i == m {
	//		if isNum {
	//			return 1
	//		}
	//		return 0
	//	}
	//
	//	if !isLimit && isNum && memo[i][mask] != -1 {
	//		return memo[i][mask]
	//	}
	//
	//	res := 0
	//
	//	if !isNum {
	//		res += dfs(i+1, mask, false, false)
	//	}
	//	up := 9
	//	if isLimit {
	//		up = int(s[i] - '0')
	//	}
	//	low := 0
	//	if !isNum {
	//		low = 1
	//	}
	//	for j := low; j <= up; j++ {
	//		if mask>>j&1 == 0 {
	//			res += dfs(i+1, mask|1<<j, isLimit && j == up, true)
	//		}
	//	}
	//	if isNum && !isLimit {
	//		memo[i][mask] = res
	//	}
	//	return res
	//}

	return 1

}
