package main

import "fmt"

/*
*给你一个 n 个节点的 有向图 ，节点编号为 0 到 n - 1 ，其中每个节点 至多 有一条出边。

图用一个大小为 n 下标从 0 开始的数组 edges 表示，节点 i 到节点 edges[i] 之间有一条有向边。如果节点 i 没有出边，
那么 edges[i] == -1 。

请你返回图中的 最长 环，如果没有任何环，请返回 -1 。

一个环指的是起点和终点是 同一个 节点的路径。
*/
func longestCycle(edges []int) int {

	n := len(edges)
	inDegree := make([]int, n)
	for _, v := range edges {
		if v >= 0 {
			inDegree[v]++
		}

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

		if y >= 0 {
			inDegree[y]--
			if inDegree[y] == 0 {
				queue = append(queue, y)
			}
		}
	}
	ans := -1

	for x, v := range inDegree {

		if v == 0 {
			continue
		}
		cnt := 1
		inDegree[x] = 0
		for c := edges[x]; c != x; c = edges[c] {
			cnt++
			inDegree[c] = 0
		}
		ans = max(ans, cnt)
	}
	return ans

}

/*
*

时间戳来实现
*/
func longestCycle2(edges []int) int {
	time := make([]int, len(edges))
	clock, ans := 1, -1
	for x, t := range time {
		if t > 0 {
			continue
		}
		for startTime := clock; x >= 0; x = edges[x] {
			if time[x] > 0 { // 重复访问
				if time[x] >= startTime { // 找到了一个新的环
					ans = max(ans, clock-time[x])
				}
				break
			}
			time[x] = clock
			clock++
		}
	}
	return ans
}

func main() {
	fmt.Println(longestCycle([]int{-1, 4, -1, 2, 0, 4}))
}
