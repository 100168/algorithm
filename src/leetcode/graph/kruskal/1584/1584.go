package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。

连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。

请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
*/

type pair struct {
	x int
	y int
	d int
}

func minCostConnectPoints(points [][]int) int {

	n := len(points)
	var e []pair
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := dis(points[i], points[j])
			e = append(e, pair{i, j, d})
		}
	}

	sort.Slice(e, func(i, j int) bool { return e[i].d < e[j].d })

	uf := new(unionFind)
	uf.parent = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}

	ans := 0
	cnt := n - 1

	for _, v := range e {
		if uf.union(v.x, v.y) {
			cnt--
			ans += v.d
		}
		if cnt == 0 {
			break
		}
	}
	return ans
}

func dis(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

type unionFind struct {
	parent []int
}

func (uf unionFind) find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}
func (uf unionFind) union(a, b int) bool {
	fa := uf.find(a)
	fb := uf.find(b)
	if fa == fb {
		return false
	}
	uf.parent[fa] = fb
	return true
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(minCostConnectPoints([][]int{{20, -19}, {-2, 20}, {-10, -12}, {-1, -2}, {-6, -5}, {10, -13}, {-7, -1}, {-17, -13}, {16, 14}, {13, 9}}))
}
