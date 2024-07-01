package main

import "fmt"

/*
*
力扣数据中心有 n 台服务器，分别按从 0 到 n-1 的方式进行了编号。
它们之间以 服务器到服务器 的形式相互连接组成了一个内部集群，连接是无向的。
用  connections 表示集群网络，connections[i] = [a, b] 表示服务器 a 和 b 之间形成连接。
任何服务器都可以直接或者间接地通过网络到达任何其他服务器。

关键连接 是在该集群中的重要连接，假如我们将它移除，便会导致某些服务器无法访问其他服务器。

请你以任意顺序返回该集群内的所有 关键连接 。

1.割点: 删掉这个点和这个点有关的边,图就不是连通图,分裂成为了多个不相连的子图
2.割边: 删除这条边后,图就不是连通图,分裂成为多个不相连的子图
*/
func criticalConnections(n int, connections [][]int) [][]int {

	g := make([][]int, n)
	for _, v := range connections {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	t := 0
	dfn := make([]int, n)
	low := make([]int, n)

	ans := make([][]int, 0)
	var dfs func(int, int) int

	dfs = func(x int, fa int) int {

		t++
		dfn[x] = t
		low[x] = t
		for _, v := range g[x] {
			if dfn[v] == 0 {
				next := dfs(v, x)
				if next > dfn[x] {
					ans = append(ans, []int{x, v})
				}
				low[x] = min(low[x], next)
			} else if v != fa {
				low[x] = min(low[x], dfn[v])
			}
		}
		return low[x]
	}
	dfs(0, -1)

	return ans
}

func main() {
	fmt.Println(criticalConnections(4, [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
}
