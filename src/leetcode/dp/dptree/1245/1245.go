package main

func treeDiameter(edges [][]int) int {

	n := len(edges)
	g := make([][]int, n+1)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := 0

	var dfs func(int, int) int

	dfs = func(x, fa int) int {

		curMax := 0

		for _, y := range g[x] {
			if y == fa {
				continue
			}
			l := dfs(y, x) + 1
			ans = max(ans, curMax+l)
			curMax = max(curMax, l)
		}
		return curMax
	}
	dfs(0, -1)
	return ans
}
