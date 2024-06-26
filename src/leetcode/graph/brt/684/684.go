package main

import "fmt"

/*
*树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间
，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。
*/
func findRedundantConnection(edges [][]int) []int {

	n := len(edges)
	g := make([][]int, n)

	inDeg := make([]int, n)
	for _, v := range edges {
		x, y := v[0]-1, v[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		inDeg[x]++
		inDeg[y]++
	}

	var queue []int
	for i := range inDeg {
		if inDeg[i] == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, v := range g[cur] {
			inDeg[v]--
			if inDeg[v] == 1 {
				queue = append(queue, v)
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		x, y := edges[i][0]-1, edges[i][1]-1

		if inDeg[x] == 2 && inDeg[y] == 2 {
			return edges[i]
		}
	}
	return nil
}

func main() {

	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}
