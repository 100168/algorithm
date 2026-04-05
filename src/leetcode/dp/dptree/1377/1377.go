package main

func frogPosition(n int, edges [][]int, t int, target int) (ans float64) {
	g := make([][]int, n+1)

	g[1] = append(g[1], 0)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int, int, int) bool

	dfs = func(x int, fa int, leftT int, prod int) bool {

		if x == target && (leftT == 0 || len(g[x]) == 1) {
			ans = 1 / float64(prod)
			return true
		}
		if x == target || leftT == 0 {
			return false
		}
		for _, y := range g[x] {
			if y != fa && dfs(y, x, leftT-1, prod*(len(g[x])-1)) {
				return true
			}
		}
		return false
	}
	dfs(1, 0, t, 1)
	return ans
}
