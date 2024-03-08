//给你一个有 n 个节点的 有向带权 图，节点编号为 0 到 n - 1 。图中的初始边用数组 edges 表示，其中 edges[i] = [fromi,
//toi, edgeCosti] 表示从 fromi 到 toi 有一条代价为 edgeCosti 的边。
//
// 请你实现一个 Graph 类：
//
//
// Graph(int n, int[][] edges) 初始化图有 n 个节点，并输入初始边。
// addEdge(int[] edge) 向边集中添加一条边，其中 edge = [from, to, edgeCost] 。数据保证添加这条边之前对应的两
//个节点之间没有有向边。
// int shortestPath(int node1, int node2) 返回从节点 node1 到 node2 的路径 最小 代价。如果路径不存在，
//返回 -1 。一条路径的代价是路径中所有边代价之和。
//
//
//
//
// 示例 1：
//
//
//
// 输入：
//["Graph", "shortestPath", "shortestPath", "addEdge", "shortestPath"]
//[[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]
//], [0, 3]]
//输出：
//[null, 6, -1, null, 6]
//
//解释：
//Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
//g.shortestPath(3, 2); // 返回 6 。从 3 到 2 的最短路径如第一幅图所示：3 -> 0 -> 1 -> 2 ，总代价为 3 +
// 2 + 1 = 6 。
//g.shortestPath(0, 3); // 返回 -1 。没有从 0 到 3 的路径。
//g.addEdge([1, 3, 4]); // 添加一条节点 1 到节点 3 的边，得到第二幅图。
//g.shortestPath(0, 3); // 返回 6 。从 0 到 3 的最短路径为 0 -> 1 -> 3 ，总代价为 2 + 4 = 6 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= edges.length <= n * (n - 1)
// edges[i].length == edge.length == 3
// 0 <= fromi, toi, from, to, node1, node2 <= n - 1
// 1 <= edgeCosti, edgeCost <= 10⁶
// 图中任何时候都不会有重边和自环。
// 调用 addEdge 至多 100 次。
// 调用 shortestPath 至多 100 次。
//
//
// Related Topics 图 设计 最短路 堆（优先队列） 👍 16 👎 0

package main

import (
	"container/heap"
	"math"
)

type pair struct {
	x int
	d int
}

type hp []pair

func (h *hp) Less(i, j int) bool {

	return (*h)[i].d < (*h)[j].d
}
func (h *hp) Swap(i, j int) {

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

type edge struct {
	to int
	d  int
}

type Graph struct {
	g    [][]edge
	size int
}

func Constructor(n int, edges [][]int) Graph {
	g := new(Graph)
	g.g = make([][]edge, n)
	for _, e := range edges {
		x, y, d := e[0], e[1], e[2]
		g.g[x] = append(g.g[x], edge{y, d})
	}
	g.size = n
	return *g
}

func (g *Graph) AddEdge(e []int) {
	x, y, d := e[0], e[1], e[2]
	g.g[x] = append(g.g[x], edge{y, d})

}

func (g *Graph) ShortestPath(node1 int, node2 int) int {

	dis := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		dis[i] = math.MaxInt
	}
	dis[node1] = 0
	hp := &hp{{node1, 0}}
	for hp.Len() > 0 {
		p := heap.Pop(hp).(pair)
		x := p.x
		if p.d > dis[x] {
			continue
		}
		for _, e := range g.g[x] {
			y := e.to
			newDis := p.d + e.d
			if newDis < dis[y] {
				dis[y] = newDis
				heap.Push(hp, pair{y, newDis})
			}
		}
	}
	if dis[node2] == math.MaxInt {
		return -1
	}
	return dis[node2]
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
//leetcode submit region end(Prohibit modification and deletion)
