package main

import (
	"slices"
	"sort"
)

// 先离散化+dp
func jobScheduling(startTime []int, endTime []int, profit []int) int {

	n := len(startTime)
	nums := make([]int, 0)

	for i := 0; i < n; i++ {
		nums = append(nums, startTime[i])
		nums = append(nums, endTime[i])
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
		start, end, gold := rkMap[startTime[i]], rkMap[endTime[i]], profit[i]

		groups[end] = append(groups[end], pair{start, gold})
	}

	f := make([]int, m+1)
	for end, g := range groups {
		f[end+1] = f[end]
		for _, p := range g {
			f[end+1] = max(f[end+1], f[p.start+1]+p.gold)
		}
	}
	return f[m]
}
