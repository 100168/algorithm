package main

import "math"

var amountBob []int

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	amountBob = make([]int, n)
	for i := range amountBob {
		amountBob[i] = n
	}
	g := make([][]int, n)
	g[0] = append(g[0], -1)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	dfsBob(bob, -1, 0, g, amount)
	return dfs(0, -1, 0, g, amount)

}

func dfsBob(x, fa, t int, g [][]int, amount []int) bool {

	if x == 0 {
		amountBob[0] = t
		return true
	}

	for _, y := range g[x] {

		if y != fa && dfsBob(y, x, t+1, g, amount) {
			amountBob[x] = t
			return true
		}
	}
	return false

}

func dfs(x, f, t int, g [][]int, amount []int) int {

	ans := 0
	if amountBob[x] == t {
		ans += amount[x] / 2
	} else if amountBob[x] > t {
		ans += amount[x]
	}

	curMax := math.MinInt
	for _, y := range g[x] {
		if y != f {
			curMax = max(curMax, dfs(y, x, t+1, g, amount))
		}
	}
	if len(g[x]) == 1 {
		curMax = 0
	}
	return ans + curMax
}

func main() {
	println(mostProfitablePath([][]int{{0, 1}}, 1, []int{-7280, 2350}))
}
