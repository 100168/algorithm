package main

import (
	"fmt"
	"sort"
)

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {

	n := len(edges1) + 1
	m := len(edges2) + 1

	g1 := make([][]int, n)
	g2 := make([][]int, m)

	for _, v := range edges1 {
		x, y := v[0], v[1]

		g1[x] = append(g1[x], y)
		g1[y] = append(g1[y], x)
	}
	for _, v := range edges2 {
		x, y := v[0], v[1]

		g2[x] = append(g2[x], y)
		g2[y] = append(g2[y], x)
	}

	ans := 0
	deep1 := make([]int, n)
	deep2 := make([]int, m)
	var dfs func(int, int, int, int, [][]int)

	dfs = func(x, fa, deep int, t int, g [][]int) {
		if deep <= t {
			ans++
		}
		for _, v := range g[x] {
			if v == fa {
				continue
			}
			dfs(v, x, deep+1, t, g)
		}
	}

	for i := 0; i < n; i++ {

		ans = 0
		dfs(i, -1, 0, k, g1)
		deep1[i] = ans

	}

	for i := 0; i < m; i++ {

		ans = 0
		dfs(i, -1, 0, k-1, g2)
		deep2[i] = ans

	}

	sort.Ints(deep2)

	cc := make([]int, n)

	for i := 0; i < n; i++ {

		cc[i] += deep1[i] + deep2[len(deep2)-1]
	}

	return cc

}

func main() {
	//fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}, 2))
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, [][]int{{0, 1}, {1, 2}, {2, 3}}, 1))
}
