package main

func isBipartite(graph [][]int) bool {

	n := len(graph)

	colors := make([]int, n)
	var dfs func(int, int) bool

	dfs = func(x, c int) bool {

		colors[x] = c

		for _, y := range graph[x] {
			if colors[y] == 0 {
				if !dfs(y, 3^c) {
					return false
				}
			} else {
				if colors[y] == c {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if colors[i] != 0 {
			continue
		}
		if !dfs(i, 1) {
			return false
		}
	}
	return true

}
