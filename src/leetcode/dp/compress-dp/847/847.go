package main

import (
	"fmt"
	"math"
)

/*
*
存在一个由 n 个节点组成的无向连通图，图中的节点按从 0 到 n - 1 编号。

给你一个数组 graph 表示这个图。其中，graph[i] 是一个列表，由所有与节点 i 直接相连的节点组成。

返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。
*/
func shortestPathLength(graph [][]int) int {

	n := len(graph)

	memo := make([][]int, 1<<n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	memo2 := make([][][]int, n)

	for i := range memo2 {
		memo2[i] = make([][]int, n)
		for j := range memo2[i] {
			memo2[i][j] = make([]int, n)
			for k := range memo2[i][j] {
				memo2[i][j][k] = -1
			}
		}
	}
	var floyd func(int, int, int) int

	floyd = func(k int, i int, j int) int {
		if k < 0 {
			for _, v := range graph[i] {
				if v == j {
					return 1
				}
			}
			return math.MaxInt / 2
		}
		if memo2[k][i][j] != -1 {
			return memo2[k][i][j]
		}
		memo2[k][i][j] = min(floyd(k-1, i, j), floyd(k-1, i, k)+floyd(k-1, k, j))
		return memo2[k][i][j]

	}

	var dfs func(int, int) int

	dfs = func(mask int, pre int) int {

		if mask == 1<<n-1 {
			return 0
		}

		if memo[mask][pre] != -1 {
			return memo[mask][pre]
		}

		cur := math.MaxInt / 2
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				cur = min(cur, dfs(1<<j|mask, j)+floyd(n-1, j, pre))
			}
		}
		memo[mask][pre] = cur
		return cur
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dfs(1<<i, i))
	}
	return ans
}

func shortestPathLength2(graph [][]int) int {

	n := len(graph)

	memo := make([][]int, 1<<n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2
		}
	}
	for x, v := range graph {
		for _, y := range v {
			dp[x][y] = 1
		}
	}

	for k := range dp {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
			}
		}
	}

	var dfs func(int, int) int

	dfs = func(mask int, pre int) int {

		if mask == 1<<n-1 {
			return 0
		}

		if memo[mask][pre] != -1 {
			return memo[mask][pre]
		}

		cur := math.MaxInt / 2
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				cur = min(cur, dfs(1<<j|mask, j)+dp[j][pre])
			}
		}
		memo[mask][pre] = cur
		return cur
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dfs(1<<i, i))
	}
	return ans
}
func main() {
	fmt.Println(shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}}))
	fmt.Println(shortestPathLength2([][]int{{1, 2, 3}, {0}, {0}, {0}}))
}
