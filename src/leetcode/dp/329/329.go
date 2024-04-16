package main

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)

		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if cache[x][y] != -1 {
			return cache[x][y]
		}
		res := 1
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && matrix[nx][ny] > matrix[x][y] {
				res = max(res, dfs(nx, ny)+1)
			}
		}
		cache[x][y] = res
		return res
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}

	}

	return ans

}

func main() {

	println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))

}
