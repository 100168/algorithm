package main

import (
	"fmt"
)

/*
*
有 n 座城市，编号从 1 到 n 。编号为 x 和 y 的两座城市直接连通的前提是： x 和 y 的公因数中，
至少有一个 严格大于 某个阈值 threshold 。更正式地说，如果存在整数 z ，且满足以下所有条件，则编号 x 和 y 的城市之间有一条道路：

x % z == 0
y % z == 0
z > threshold
给你两个整数 n 和 threshold ，以及一个待查询数组，请你判断每个查询 queries[i] = [ai, bi]
指向的城市 ai 和 bi 是否连通（即，它们之间是否存在一条路径）。

返回数组 answer ，其中answer.length == queries.length 。
如果第 i 个查询中指向的城市 ai 和 bi 连通，则 answer[i] 为 true ；如果不连通，则 answer[i] 为 false 。
*/
func areConnected(n int, threshold int, queries [][]int) []bool {

	parent := make([]int, n+1)

	for i := range parent {
		parent[i] = i
	}

	find := func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) {
		fa := find(a)
		fb := find(b)
		if fa == fb {
			return
		}
		parent[fa] = fb
	}

	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}

	for i := threshold + 1; i <= n; i++ {
		if primes[i] {
			for j := i; j <= n; j += i {
				union(i, j)
				primes[j] = false
			}
		}
	}
	ans := make([]bool, len(queries))

	for i, q := range queries {
		ans[i] = find(q[0]) == find(q[1])
	}
	return ans
}
func areConnected2(n int, threshold int, queries [][]int) []bool {

	parent := make([]int, n+1)

	for i := range parent {
		parent[i] = i
	}

	find := func(a int) int {
		for parent[a] != a {
			parent[a] = parent[parent[a]]
			a = parent[a]
		}
		return a
	}

	union := func(a, b int) {
		fa := find(a)
		fb := find(b)
		if fa == fb {
			return
		}
		parent[fa] = fb
	}
	for i := threshold + 1; i <= n; i++ {
		first := i
		for j := i; j <= n; j += first {
			union(first, j)
		}
	}
	ans := make([]bool, len(queries))

	for i, q := range queries {
		ans[i] = find(q[0]) == find(q[1])
	}
	return ans
}

func main() {
	fmt.Println(areConnected(6, 0, [][]int{{4, 5}, {3, 4}, {3, 2}, {2, 6}, {1, 3}}))
}
