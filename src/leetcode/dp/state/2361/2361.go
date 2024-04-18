package main

import "math"

func minimumCosts(regular []int, express []int, expressCost int) []int64 {
	n := len(regular)
	r := make([]int, n+1)
	e := make([]int, n+1)
	//起点在普通车站0位置
	e[0] = math.MaxInt / 2

	cost := make([]int64, n)
	cost[0] = min(int64(r[0]), int64(e[0]))
	for i := 1; i <= n; i++ {
		//必须先计算特快
		e[i] = min(e[i-1]+express[i-1], r[i-1]+expressCost+express[i-1])
		r[i] = min(r[i-1]+regular[i-1], e[i])
		cost[i-1] = min(int64(r[i]), int64(e[i]))
	}

	return cost
}
