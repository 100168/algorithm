package main

func minEdgeReversals(n int, edges [][]int) []int {

	g := make([][]int, n)

	type pair struct{ x, y int }

	dirs := make(map[pair]bool)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		dirs[pair{x, y}] = true

	}

	ans := make([]int, n)
	var dfs func(int, int)

	dfs = func(x, fa int) {

		for _, y := range g[x] {

			if y == fa {
				continue
			}
			if dirs[pair{y, x}] {
				ans[0]++
			}
			dfs(y, x)
		}
	}

	dfs(0, -1)

	var reRoot func(int, int)

	reRoot = func(x, fa int) {

		for _, y := range g[x] {

			if y == fa {
				continue
			}
			cur := ans[x]
			if dirs[pair{x, y}] {
				cur++
			} else {
				cur--
			}
			ans[y] = cur
			reRoot(y, x)
		}
	}
	reRoot(0, -1)
	return ans
}
