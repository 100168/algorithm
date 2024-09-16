package main

/**
给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。

例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。

求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。



示例 1：

输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
输出：2
解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
示例 2：

输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
输出：-1
*/

func numBusesToDestination(routes [][]int, source int, target int) int {

	if source == target {
		return 0
	}
	g := make(map[int][]int)

	for i, route := range routes {

		for _, v := range route {
			g[v] = append(g[v], i)
		}
	}

	var q []int

	visited := make(map[int]bool)

	visitedBus := make(map[int]bool)
	for _, v := range g[source] {

		if visitedBus[v] {
			continue
		}
		for _, v := range routes[v] {
			if !visited[v] {
				q = append(q, v)
			}
			visited[v] = true
		}
		visitedBus[v] = true
	}

	ans := 1
	for len(q) > 0 {
		temp := q
		q = nil
		for _, v := range temp {
			if v == target {
				return ans
			}
			for _, w := range g[v] {
				if visitedBus[w] {
					continue
				}
				visitedBus[w] = true
				for _, v := range routes[w] {
					if !visited[v] {
						q = append(q, v)
						visited[v] = true
					}
				}
			}
		}
		ans++
	}
	return -1
}
