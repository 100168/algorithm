package main

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {

	s := 0

	for _, v := range values {
		s += v
	}
	n := len(edges) + 1
	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int

	dfs = func(x, fa int) int {

		s := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			s += dfs(y, x)
		}
		if s == 0 {
			return values[x]
		}

		return min(s, values[x])
	}
	return int64(s) - int64(dfs(0, -1))
}
