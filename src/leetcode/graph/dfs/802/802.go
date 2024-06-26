package main

import "fmt"

/*
*
有一个有 n 个节点的有向图，节点按 0 到 n - 1 编号。图由一个 索引从 0 开始 的 2D 整数数组 graph表示，
graph[i]是与节点 i 相邻的节点的整数数组，这意味着从节点 i 到 graph[i]中的每个节点都有一条边。

如果一个节点没有连出的有向边，则该节点是 终端节点 。如果从该节点开始的所有可能路径都通向 终端节点 ，则该节点为 安全节点 。

返回一个由图中所有 安全节点 组成的数组作为答案。答案数组中的元素应当按 升序 排列。

思路：
1.三色标记

	0表示未访问，1表示正在访问或者有环，2表示可以到达端点
*/
func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var safe func(int) bool
	safe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for _, y := range graph[x] {
			if !safe(y) {
				return false
			}
		}
		color[x] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}

/*
*
top排序
*/
func eventualSafeNodes2(graph [][]int) (ans []int) {
	n := len(graph)
	inDegree := make([]int, n)

	reg := make([][]int, n)

	for i := range graph {
		for _, v := range graph[i] {
			reg[v] = append(reg[v], i)
			inDegree[i]++
		}
	}

	var queue []int

	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)

		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, v := range reg[cur] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			ans = append(ans, i)
		}
	}
	return
}

func main() {
	fmt.Println(eventualSafeNodes2([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
}
