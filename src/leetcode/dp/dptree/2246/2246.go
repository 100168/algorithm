package main

/*
给你一棵 树（即一个连通、无向、无环图），根节点是节点 0 ，这棵树由编号从 0 到 n - 1 的 n 个节点组成。
用下标从 0 开始、长度为 n 的数组 parent 来表示这棵树，其中 parent[i] 是节点 i 的父节点，由于节点 0 是根节点，所以 parent[0] == -1 。

另给你一个字符串 s ，长度也是 n ，其中 s[i] 表示分配给节点 i 的字符。

请你找出路径上任意一对相邻节点都没有分配到相同字符的 最长路径 ，并返回该路径的长度。
*/
func longestPath(parent []int, s string) int {

	ans := 0
	n := len(s)

	g := make([][]int, n)

	for x, y := range parent {
		if y == -1 {
			continue
		}
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int

	dfs = func(fa, x int) int {

		maxL := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			l := dfs(x, y) + 1
			if s[x] != s[y] {
				ans = max(ans, maxL+l)
				maxL = max(maxL, l)
			}
		}
		return maxL
	}

	dfs(-1, 0)
	return ans + 1
}

func main() {
	println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
}
