package main

import (
	"fmt"
	"math"
	"math/bits"
)

/*
*
一系列高速公路连接从 0 到 n - 1 的 n 个城市。给定一个二维整数数组 highways，
其中 highways[i] = [city1i, city2i, tolli] 表示有一条高速公路连接 city1i 和city2i，
允许一辆汽车从 city1i 前往 city2i，反之亦然，费用为 tolli。
给你一个整数 k，你要正好经过 k 条公路。你可以从任何一个城市出发，但在旅途中每个城市最多只能访问一次。
返回您旅行的最大费用。如果没有符合要求的行程，则返回 -1。
*/
func maximumCost(n int, highways [][]int, k int) int {
	g := make([]map[int]int, n)

	for i := range g {
		g[i] = make(map[int]int)
	}

	for _, v := range highways {
		x, y, c := v[0], v[1], v[2]
		g[x][y] = c
		g[y][x] = c
	}

	if k > n-1 {
		return -1
	}
	memo := make([]map[int]int, 1<<n)

	for i := range memo {
		memo[i] = make(map[int]int)

	}

	var dfs func(int, int) int

	dfs = func(mask, pre int) int {

		bitOnes := bits.OnesCount(uint(mask))
		if bitOnes-1 == k {
			return 0
		}

		if _, ok := memo[mask][pre]; ok {
			return memo[mask][pre]
		}

		cur := math.MinInt / 2
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				if _, ok := g[pre][j]; ok {
					cur = max(cur, dfs(1<<j|mask, j)+g[pre][j])
				}
			}
		}
		memo[mask][pre] = cur
		return cur
	}
	ans := math.MinInt / 2
	for i := 0; i < n; i++ {
		ans = max(ans, dfs(1<<i, i))
	}

	if ans < 0 {
		return -1
	}
	return ans
}

func main() {

	fmt.Println(maximumCost(5, [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 3}, {3, 4, 2}}, 3))
}
