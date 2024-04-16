package main

func longestString(x, y, z int) int {

	var dfs func(int, int, int, int) int

	cache := make([][][][]int, x+1)
	for i := range cache {
		cache[i] = make([][][]int, y+1)
		for j := range cache[i] {
			cache[i][j] = make([][]int, z+1)
			for k := range cache[i][j] {
				cache[i][j][k] = make([]int, 3)
				for t := range cache[i][j][k] {
					cache[i][j][k][t] = -1
				}
			}
		}
	}
	dfs = func(x, y, z, k int) int {
		cur := 0

		if cache[x][y][z][k] != -1 {
			return cache[x][y][z][k]
		}
		if k == 0 {
			if y > 0 {
				cur = dfs(x, y-1, z, 1) + 2
			}
			if z > 0 {
				cur = max(dfs(x, y, z-1, 2)+2, cur)
			}
		}
		if k == 1 {
			if x > 0 {
				cur = dfs(x-1, y, z, 0) + 2
			}
		}
		if k == 2 {
			if z > 0 {
				cur = dfs(x, y, z-1, 2) + 2
			}
			if y > 0 {
				cur = max(cur, dfs(x, y-1, z, 1)+2)
			}
		}
		cache[x][y][z][k] = cur
		return cur

	}
	return max(dfs(x, y, z, 0), dfs(x, y, z, 1), dfs(x, y, z, 2))
}
