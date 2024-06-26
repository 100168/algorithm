package main

import "fmt"

/*
*
现有一个有向图，其中包含 n 个节点，节点编号从 0 到 n - 1 。此外，该图还包含了 n 条有向边。

给你一个下标从 0 开始的数组 edges ，其中 edges[i] 表示存在一条从节点 i 到节点 edges[i] 的边。

想象在图上发生以下过程：

你从节点 x 开始，通过边访问其他节点，直到你在 此过程 中再次访问到之前已经访问过的节点。
返回数组 answer 作为答案，其中 answer[i] 表示如果从节点 i 开始执行该过程，你可以访问到的不同节点数。
*/
func countVisitedNodes(edges []int) []int {

	n := len(edges)

	inDegree := make([]int, n)

	for _, v := range edges {
		inDegree[v]++
	}
	var queue []int

	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		y := edges[x]

		if inDegree[y]--; inDegree[y] == 0 {
			queue = append(queue, y)
		}
	}

	ans := make([]int, n)

	for i, v := range inDegree {

		if v == 0 {
			continue
		}

		cnt := 1
		for c := edges[i]; c != i; c = edges[c] {
			cnt++
			inDegree[c] = 0
		}

		for c := edges[i]; c != i; c = edges[c] {
			ans[c] = cnt
		}
		ans[i] = cnt

	}

	var dfs func(int) int

	dfs = func(x int) int {

		if ans[x] != 0 {
			return ans[x]
		}

		return dfs(edges[x]) + 1
	}

	for i := range ans {
		if ans[i] == 0 {
			ans[i] = dfs(edges[i]) + 1
		}
	}
	return ans

}

func main() {
	fmt.Println(countVisitedNodes([]int{1, 2, 0, 0}))
}
