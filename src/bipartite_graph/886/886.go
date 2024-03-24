package main

func possibleBiPartition(n int, dislikes [][]int) bool {

	colors := make([]int, n+1)

	var dfs func(int, int) bool

	g := make([][]int, n+1)
	for _, v := range dislikes {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	dfs = func(x, color int) bool {
		colors[x] = color
		for _, y := range g[x] {
			if colors[y] == 0 && !dfs(y, color^3) || colors[y] == color {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true

}
