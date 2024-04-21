package main

import (
	"slices"
	"sort"
)

// 离散化+dp
func maxValue(events [][]int, k int) int {
	n := len(events)
	nums := make([]int, 0)

	for i := 0; i < n; i++ {
		nums = append(nums, events[i][0])
		nums = append(nums, events[i][1])
	}
	sort.Ints(nums)
	nums = slices.Compact(nums)

	rkMap := make(map[int]int)

	for i, v := range nums {
		rkMap[v] = i
	}
	m := len(nums)
	type pair struct{ start, gold int }
	groups := make([][]pair, m)
	for i := 0; i < n; i++ {
		start, end, gold := rkMap[events[i][0]], rkMap[events[i][1]], events[i][2]
		groups[end] = append(groups[end], pair{start, gold})
	}

	f := make([][]int, m+1)

	for i := range f {
		f[i] = make([]int, k+1)
	}
	for end, g := range groups {
		for j := 1; j <= k; j++ {
			f[end+1][j] = f[end][j]
			for _, p := range g {
				f[end+1][j] = max(f[end+1][j], f[p.start][j-1]+p.gold)
			}
		}

	}
	return f[m][k]
}
