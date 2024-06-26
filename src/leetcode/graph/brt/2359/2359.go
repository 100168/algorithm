package main

import (
	"fmt"
	"math"
)

/*
*
给你一个 n 个节点的 有向图 ，节点编号为 0 到 n - 1 ，每个节点 至多 有一条出边。

有向图用大小为 n 下标从 0 开始的数组 edges 表示，表示节点 i 有一条有向边指向 edges[i] 。
如果节点 i 没有出边，那么 edges[i] == -1 。

同时给你两个节点 node1 和 node2 。

请你返回一个从 node1 和 node2 都能到达节点的编号，使节点 node1 和节点 node2 到这个节点的距离 较大值最小化。
如果有多个答案，请返回 最小 的节点编号。如果答案不存在，返回 -1 。

注意 edges 可能包含环。
*/
func closestMeetingNode(edges []int, node1 int, node2 int) int {

	n := len(edges)
	deep1 := make([]int, n)

	for i := range deep1 {
		deep1[i] = -1
	}

	deep2 := make([]int, n)

	for i := range deep2 {
		deep2[i] = -1
	}

	d := 0
	for x := node1; x >= 0 && deep1[x] == -1; x = edges[x] {
		deep1[x] = d
		d++
	}
	d = 0
	for x := node2; x >= 0 && deep2[x] == -1; x = edges[x] {
		deep2[x] = d
		d++
	}

	ans := -1
	size := math.MaxInt
	for i := 0; i < n; i++ {
		if deep1[i] >= 0 && deep2[i] >= 0 {
			cur := max(deep1[i], deep2[i])
			if cur < size || cur == size && i < ans {
				ans = i
				size = cur
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
}
