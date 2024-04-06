package main

func hasValidPath(grid [][]byte) bool {

	m := len(grid)
	n := len(grid[0])

	var dfs func(int, int, int) bool

	dfs = func(x, y, cnt int) bool {

		if cnt < 0 {
			return false
		}
		if x == 0 && y == 0 {
			if grid[x][y] == '(' && cnt == 1 {
				return true
			}
			return false
		}
		if grid[x][y] == '(' {
			cnt--
		} else {
			cnt++
		}
		cur := false
		if x-1 >= 0 {
			cur = cur || dfs(x-1, y, cnt)
		}
		if y-1 > 0 {
			cur = cur || dfs(x, y-1, cnt)
		}
		return cur

	}

	return dfs(m-1, n-1, 0)

}
