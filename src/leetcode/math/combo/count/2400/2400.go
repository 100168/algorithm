package main

import "fmt"

/*
*
给你两个 正 整数 startPos 和 endPos 。最初，你站在 无限 数轴上位置 startPos 处。
在一步移动中，你可以向左或者向右移动一个位置。

给你一个正整数 k ，返回从 startPos 出发、恰好 移动 k 步并到达 endPos 的 不同 方法数目。
由于答案可能会很大，返回对 109 + 7 取余 的结果。

如果所执行移动的顺序不完全相同，则认为两种方法不同。

注意：数轴包含负整数。
*/
func numberOfWays(startPos int, endPos int, k int) int {

	mod := int(1e9 + 7)
	if abs(startPos-endPos) > k {
		return 0
	}

	dp := make([]map[int]int, k+1)

	for i := range dp {
		dp[i] = make(map[int]int)
	}

	var dfs func(int, int) int

	dfs = func(s int, k int) int {

		if k < 0 {
			return 0
		}
		if s == endPos && k == 0 {
			return 1
		}
		if _, ok := dp[k][s]; ok {
			return dp[k][s]
		}
		cur := (dfs(s-1, k-1) + dfs(s+1, k-1)) % mod
		dp[k][s] = cur
		return cur
	}

	return dfs(startPos, k)
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(numberOfWays(1, 2, 3))
}
