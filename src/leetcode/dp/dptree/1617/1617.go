package main

import (
	"fmt"
	"math/bits"
)

/*
*
给你 n 个城市，编号为从 1 到 n 。
同时给你一个大小为 n-1 的数组 edges ，其中 edges[i] = [ui, vi] 表示城市 ui 和 vi 之间有一条双向边。
题目保证任意城市之间只有唯一的一条路径。换句话说，所有城市形成了一棵 树 。

一棵 子树 是城市的一个子集，且子集中任意城市之间可以通过子集中的其他城市和边到达。
两个子树被认为不一样的条件是至少有一个城市在其中一棵子树中存在，但在另一棵子树中不存在。

对于 d 从 1 到 n-1 ，请你找到城市间 最大距离 恰好为 d 的所有子树数目。

请你返回一个大小为 n-1 的数组，其中第 d 个元素（下标从 1 开始）是城市间 最大距离 恰好等于 d 的子树数目。

请注意，两个城市间距离定义为它们之间需要经过的边的数目。
*/
func countSubgraphsForEachDiameter(n int, edges [][]int) []int {

	g := make([][]int, n)

	ans := make([]int, n-1)
	for _, v := range edges {
		x, y := v[0]-1, v[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	for i := 0; i < 1<<n; i++ {

		if i&(i-1) == 0 {
			continue
		}
		first := bits.TrailingZeros(uint(i))

		maxL := 0
		var dfs func(int, int) int

		vis := i ^ 1<<first

		dfs = func(x, fa int) int {
			cur := 0
			for _, y := range g[x] {
				if y == fa || vis>>y&1 == 0 {
					continue
				}
				vis ^= 1 << y
				l := dfs(y, x) + 1
				maxL = max(maxL, l+cur)
				cur = max(l, cur)
			}
			return cur
		}
		dfs(first, -1)

		if vis == 0 {
			ans[maxL-1]++
		}

	}
	return ans
}

func main() {
	fmt.Println(countSubgraphsForEachDiameter(4, [][]int{{1, 2}, {2, 3}, {2, 4}}))
}
