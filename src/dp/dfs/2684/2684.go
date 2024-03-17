package main

import "fmt"

func maxMoves(grid [][]int) int {

	m := len(grid)

	n := len(grid[0])

	ans := 0

	var dfs func(int, int)

	dfs = func(x, y int) {
		ans = max(ans, y)
		if ans == n-1 {
			return
		}
		for k := max(x-1, 0); k < min(x+2, m); k++ {
			if grid[k][y+1] > grid[x][y] {
				dfs(k, y+1)
			}
		}
		grid[x][y] = 0

	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
	}
	return ans

}

func maxMoves2(grid [][]int) int {

	m := len(grid)

	n := len(grid[0])

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(x, y int) int {
		if dp[x][y] != -1 {
			return dp[x][y]
		}
		cur := 0
		for i := max(x-1, 0); i < min(x+2, m) && y+1 < n; i++ {
			if grid[x][y] < grid[i][y+1] {
				cur = max(dfs(i, y+1)+1, cur)
			}
		}
		dp[x][y] = cur
		return cur
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = max(dfs(i, 0), ans)
	}
	return ans

}

func main() {
	fmt.Println(maxMoves([][]int{{1, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
}
