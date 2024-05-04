package main

import "fmt"

func maximumPoints(edges [][]int, coins []int, k int) int {

	n := len(edges) + 1

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]

		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([]map[int]int, n)

	for i := range memo {
		memo[i] = make(map[int]int)
	}
	var dfs func(int, int, int) int

	dfs = func(x, fa int, t int) int {

		if 1<<t > 10000 {
			return 0
		}
		if _, ok := memo[x][t]; ok {
			return memo[x][t]
		}

		coin := int(coins[x] / (1 << t))
		f0, f1 := coin-k, coin/2
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			f0 += dfs(y, x, t)
			f1 += dfs(y, x, t+1)
		}
		memo[x][t] = max(f0, f1)
		return max(f0, f1)
	}
	return dfs(0, -1, 0)
}

func main() {
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}, []int{10, 10, 3, 20, 10}, 5))
}
