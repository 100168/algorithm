package main

import (
	"fmt"
)

/*
*

思路：
1.先从上往下计算每个节点的深度
*/
func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {

	n := len(edges1) + 1
	m := len(edges2) + 1

	deep1 := getDeep(edges1, n)
	deep2 := getDeep(edges2, m)

	s1 := 0

	for _, v := range deep1 {
		s1 += v
	}

	for i, v := range deep1 {
		if v == 0 {
			deep1[i] = n - s1
		} else {
			deep1[i] = s1
		}
	}
	s2 := 0

	for _, v := range deep2 {
		s2 += v
	}

	s2 = max(s2, m-s2)

	ans := make([]int, n)

	for i := range ans {
		ans[i] = deep1[i] + s2
	}
	return ans
}

func getDeep(e [][]int, n int) []int {

	g := make([][]int, n)

	for _, v := range e {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)

	var dfs func(x, f, d int)

	dfs = func(x, f, d int) {

		if d%2 == 1 {
			ans[x] = 1
		}
		for _, v := range g[x] {

			if v == f {
				continue
			}
			dfs(v, x, d+1)
		}
	}

	dfs(0, -1, 0)
	return ans

}

func main() {
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}))
	fmt.Println(maxTargetNodes([][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, [][]int{{0, 1}, {1, 2}, {2, 3}}))
	fmt.Println(maxTargetNodes([][]int{{0, 1}}, [][]int{{0, 1}, {1, 2}}))
}
