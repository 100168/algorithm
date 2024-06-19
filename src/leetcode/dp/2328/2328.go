package main

/*
*
给你一个 m x n 的整数网格图 grid ，你可以从一个格子移动到 4 个方向相邻的任意一个格子。

请你返回在网格图中从 任意 格子出发，达到 任意 格子，
且路径中的数字是 严格递增 的路径数目。由于答案可能会很大，请将结果对 109 + 7 取余 后返回。

如果两条路径中访问过的格子不是完全相同的，那么它们视为两条不同的路径。
*/
func countPaths(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	mod := 1_000_000_007
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var dfs func(int, int) int

	dfs = func(x, y int) int {
		if cache[x][y] != -1 {
			return cache[x][y]
		}

		cur := 1
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] > grid[x][y] {
				cur += dfs(nx, ny)
				cur %= mod
			}
		}
		cache[x][y] = cur
		return cur
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += dfs(i, j)
			ans %= mod
		}
	}

	return ans

}

func main() {
	println(countPaths([][]int{{1, 1}, {3, 4}}))
}
