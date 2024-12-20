package main

/*
给你一个 n 个节点的无向无根图，节点编号为 0 到 n - 1 。
给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间有一条边。

每个节点都有一个价值。给你一个整数数组 price ，其中 price[i] 是第 i 个节点的价值。

一条路径的 价值和 是这条路径上所有节点的价值之和。

你可以选择树中任意一个节点作为根节点 root 。选择 root 为根的 开销 是以 root 为起点的所有路径中，价值和 最大的一条路径与最小的一条路径的差值。

请你返回所有节点作为根节点的选择中，最大 的 开销 为多少。
思路：
1.问题转换：选单前节点为根和不选单前节点为根
2.所以需要返回两个值，m1(含叶子节点最大值),m2（不含叶子节点最大值）
3.ans = max(ans,m1+s2,m2+s1)
*/
func maxOutput(n int, edges [][]int, price []int) int64 {

	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := 0

	var dfs func(int, int) (int, int)

	dfs = func(x, fa int) (int, int) {
		p := price[x]
		//m2必须是0，因为可能只有一个节点
		m1, m2 := p, 0
		for _, y := range g[x] {
			if y == fa {

				continue
			}
			// 前面最大带叶子的路径和 + 当前不带叶子的路径和
			// 前面最大不带叶子的路径和 + 当前带叶子的路径和
			s1, s2 := dfs(y, x)
			ans = max(ans, m1+s2, m2+s1)
			m1 = max(m1, s1+p)
			m2 = max(m2, s2+p)
		}
		return m1, m2
	}

	dfs(0, -1)
	return int64(ans)
}

func main() {

}
