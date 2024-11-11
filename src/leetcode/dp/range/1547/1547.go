package main

import (
	"fmt"
	"math"
	"slices"
)

/*
有一根长度为 n 个单位的木棍，棍上从 0 到 n 标记了若干位置。例如，长度为 6 的棍子可以标记如下：

给你一个整数数组 cuts ，其中 cuts[i] 表示你需要将棍子切开的位置。

你可以按顺序完成切割，也可以根据需要更改切割的顺序。

每次切割的成本都是当前要切割的棍子的长度，切棍子的总成本是历次切割成本的总和。
对棍子进行切割将会把一根木棍分成两根较小的木棍（这两根木棍的长度和就是切割前木棍的长度）。请参阅第一个示例以获得更直观的解释。

返回切棍子的 最小总成本 。
*/
func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	slices.Sort(cuts)

	m := len(cuts)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i+1 == j { // 无需切割
			return 0
		}
		p := &memo[i][j]
		if *p > 0 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		for k := i + 1; k < j; k++ {
			res = min(res, dfs(i, k)+dfs(k, j))
		}
		*p = res + cuts[j] - cuts[i] // 记忆化
		return *p
	}
	return dfs(0, m-1)
}

func main() {
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))
	fmt.Println(minCost(30, []int{18, 15, 13, 7, 5, 26, 25, 29}))
}
