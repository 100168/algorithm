package main

import "fmt"

/*
给你一个整数 n 和一个二维整数数组 queries。

有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。

queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。

所有查询中不会存在两个查询都满足 queries[i][0] < queries[j][0] < queries[i][1] < queries[j][1]。

返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。


*/

func shortestDistanceAfterQueries(n int, queries [][]int) []int {

	ans := make([]int, len(queries))
	s := make(map[int]int, n)
	for i := 0; i < n-1; i++ {
		s[i] = i + 1
	}
	for i, q := range queries {
		x, y := q[0], q[1]
		if _, ok := s[x]; !ok || y <= s[x] {
			ans[i] = len(s)
			continue
		}
		c := x
		for c != y {
			cc := s[c]
			delete(s, c)
			c = cc
		}
		s[x] = y
		ans[i] = len(s)

	}
	return ans
}

func main() {

	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{0, 3}, {0, 2}}))
}
