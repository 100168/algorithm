package main

import (
	"math"
	"math/bits"
)

/*
*
在由 2D 网格表示的校园里有 n 位工人（worker）和 m 辆自行车（bike），n <= m。所有工人和自行车的位置都用网格上的 2D 坐标表示。

我们为每一位工人分配一辆专属自行车，使每个工人与其分配到的自行车之间的 曼哈顿距离 最小化。

返回 每个工人与分配到的自行车之间的曼哈顿距离的最小可能总和 。

p1 和 p2 之间的 曼哈顿距离 为 Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|。
*/
func assignBikes(workers [][]int, bikes [][]int) int {

	n := len(workers)
	m := len(bikes)

	memo := make([]int, 1<<m)

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(mask int) int {

		index := bits.OnesCount(uint(mask))
		if index == n {
			return 0
		}

		if memo[mask] != -1 {
			return memo[mask]
		}

		cur := math.MaxInt
		for j := 0; j < m; j++ {
			if 1<<j&mask == 0 {
				cur = min(cur, dfs(1<<j|mask)+distance(workers[index], bikes[j]))
			}
		}
		memo[mask] = cur
		return cur
	}
	return dfs(0)
}

func distance(a, b []int) int {

	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
