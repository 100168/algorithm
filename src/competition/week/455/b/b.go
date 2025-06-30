/*
*
给你一个从 1 开始计数的整数数组 numWays，其中 numWays[i] 表示使用某些固定面值的硬币（每种面值可以使用无限次）凑出总金额 i 的方法数。每种面值都是一个正整数，并且其值最多为 numWays.length。

然而，具体的硬币面值已经丢失。你的任务是还原出可能生成这个 numWays 数组的面值集合。

返回一个按从小到大顺序排列的数组，其中包含所有可能的唯一整数面值。

如果不存在这样的集合，返回一个空数组。
*/
package main

func findCoins(numWays []int) []int {
	n := len(numWays)
	if n == 0 {
		return []int{}
	}

	aa := make([]int, n+1)
	aa[0] = 1
	copy(aa[1:], numWays)

	d := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i * 2; j <= n; j += i {
			d[j] = append(d[j], i)
		}
	}

	mm := make([]int, n+1)
	bb := make([]int, n+1)

	for m := 1; m <= n; m++ {
		sum := 0
		for k := 1; k < m; k++ {
			sum += mm[k] * aa[m-k]
		}
		mm[m] = m*aa[m] - sum

		if mm[m] < 0 {
			return []int{}
		}

		sd := 0
		for _, d := range d[m] {
			sd += d * bb[d]
		}

		mI := mm[m] - sd

		if mI == 0 {
			bb[m] = 0
		} else if mI == m {
			bb[m] = 1
		} else {
			return []int{}
		}
	}

	var S []int
	for i := 1; i <= n; i++ {
		if bb[i] == 1 {
			S = append(S, i)
		}
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for _, s := range S {
		for i := s; i <= n; i++ {
			dp[i] += dp[i-s]
		}
	}

	for i := 1; i <= n; i++ {
		if dp[i] != aa[i] {
			return []int{}
		}
	}

	return S
}
