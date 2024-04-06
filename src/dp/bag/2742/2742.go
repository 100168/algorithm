package main

import "math"

func paintWalls(cost []int, time []int) int {

	n := len(cost)
	var dfs func(int, int) int
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n*2+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if j-n > i {
			return 0
		}
		if i < 0 {
			return math.MaxInt / 2
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		cur := min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		cache[i][j] = cur
		return cur
	}

	return dfs(n-1, n)
}

func main() {
	println(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
}
