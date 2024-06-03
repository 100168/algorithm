package main

import "fmt"

/*
*
给你一个 n 个节点的树（也就是一个无环连通无向图），节点编号从 0 到 n - 1 ，
且恰好有 n - 1 条边，每个节点有一个值。树的 根节点 为 0 号点。

给你一个整数数组 nums 和一个二维数组 edges 来表示这棵树。
nums[i] 表示第 i 个点的值，edges[j] = [uj, vj] 表示节点 uj 和节点 vj 在树中有一条边。

当 gcd(x, y) == 1 ，我们称两个数 x 和 y 是 互质的 ，其中 gcd(x, y) 是 x 和 y 的 最大公约数 。

从节点 i 到 根 最短路径上的点都是节点 i 的祖先节点。一个节点 不是 它自己的祖先节点。

请你返回一个大小为 n 的数组 ans ，其中 ans[i]是离节点 i 最近的祖先节点且满足 nums[i] 和 nums[ans[i]] 是 互质的 ，
如果不存在这样的祖先节点，ans[i] 为 -1 。

思路：
1.先预处理所有互质对
2.用一个数组记录单前路径的值，和深度，深度最深的就是最近祖先
*/
const mx = 51

var coprime [mx][]int

func init() {
	// 预处理：coprime[i] 保存 [1, MX) 中与 i 互质的所有元素
	for i := 1; i < mx; i++ {
		for j := 1; j < mx; j++ {
			if gcd(i, j) == 1 {
				coprime[i] = append(coprime[i], j)
			}
		}
	}
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	type pair struct{ depth, id int }
	valDepthId := [mx]pair{}
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		val := nums[x] // x 的节点值
		// 计算与 val 互质的数中，深度最大的节点编号
		maxDepth := 0
		for _, j := range coprime[val] {
			p := valDepthId[j]
			if p.depth > maxDepth {
				maxDepth = p.depth
				ans[x] = p.id
			}
		}

		tmp := valDepthId[val]           // 用于恢复现场
		valDepthId[val] = pair{depth, x} // 保存 val 对应的节点深度和节点编号
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
			}
		}
		valDepthId[val] = tmp // 恢复现场
	}
	dfs(0, -1, 1)
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	fmt.Println(getCoprimes([]int{2, 3, 3, 2}, [][]int{{0, 1}, {1, 2}, {1, 3}}))
}
