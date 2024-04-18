package main

func knightDialer(n int) int {

	m := 4
	t := 3

	cache := make([][][]int, m)

	for i := range cache {
		cache[i] = make([][]int, t)
		for j := 0; j < t; j++ {
			cache[i][j] = make([]int, n+1)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}

	mod := 1_000_000_007

	dirs := [][]int{{1, 2}, {-1, 2}, {1, -2}, {-1, -2}, {2, 1}, {-2, 1}, {2, -1}, {-2, -1}}
	var dfs func(int, int, int) int

	dfs = func(x, y, z int) int {

		if x == 3 && (y == 0 || y == 2) {
			return 0
		}
		if x < 0 || y < 0 || x >= m || y >= t {
			return 0
		}
		if z == 0 {
			return 1
		}

		if cache[x][y][z] != -1 {
			return cache[x][y][z]
		}

		cur := 0
		for _, d := range dirs {
			cur += dfs(x+d[0], y+d[1], z-1)
			cur %= mod
		}
		cache[x][y][z] = cur
		return cur
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < t; j++ {
			ans += dfs(i, j, n-1)
			ans %= mod
		}
	}
	return ans
}

func main() {
	println(knightDialer(1))
}
