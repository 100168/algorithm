package main

import "sort"

func findOriginalArray(changed []int) []int {

	n := len(changed)

	if n%2 == 1 {
		return []int{}
	}
	ans := make([]int, 0)

	sort.Ints(changed)

	cntMap := make(map[int]int)

	for i := 0; i < n; i++ {
		cntMap[changed[i]]++
	}
	for _, v := range changed {

		if cntMap[v] == 0 {
			continue
		}
		if cntMap[v] > 0 {
			ans = append(ans, v)
		}
		if cntMap[v*2] == 0 {
			return []int{}
		}
		cntMap[v]--
		cntMap[v*2]--
	}
	return ans
}
