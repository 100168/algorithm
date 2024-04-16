package main

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
