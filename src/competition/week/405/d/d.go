package main

import (
	"index/suffixarray"
	"math"
)

/**
给你一个字符串 target、一个字符串数组 words 以及一个整数数组 costs，这两个数组长度相同。

设想一个空字符串 s。

你可以执行以下操作任意次数（包括零次）：

选择一个在范围  [0, words.length - 1] 的索引 i。
将 words[i] 追加到 s。
该操作的成本是 costs[i]。
返回使 s 等于 target 的 最小 成本。如果不可能，返回 -1。
*/

func minimumCost(target string, words []string, costs []int) int {
	minCost := map[string]uint16{}
	for i, w := range words {
		c := uint16(costs[i])
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	type pair struct{ l, cost uint16 }
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range minCost {
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
