package main

/**/

var dp []int

func countVisitedNodes(edges []int) []int {
	n := len(edges)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	for i := 0; i < n; i++ {
		dfs(i, 0, edges, make(map[int]bool))
	}
	return dp
}

func dfs(index int, curVal int, edges []int, path map[int]bool) int {

	if path[index] {
		return 1
	}
	path[index] = true
	if dp[edges[index]] > 0 {
		curVal += dp[edges[index]]
		dp[index] = curVal
	} else {
		dfs(edges[index], curVal+1, edges, path)
	}
	return dp[index]

}
