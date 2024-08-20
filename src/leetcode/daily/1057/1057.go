package main

import (
	"fmt"
	"math"
)

/*
*
在 X-Y 平面上表示的校园中，有 n 名工人和 m 辆自行车，其中 n <= m。

给定一个长度为 n 的数组 workers ，其中 worker [i] = [xi, yi] 表示第 i 个工人的位置。你也得到一个长度为 m 的自行车数组 bikers ，
其中 bikes[j] = [xj, yj] 是第 j 辆自行车的位置。所有给定的位置都是 唯一 的。

我们需要为每位工人分配一辆自行车。在所有可用的自行车和工人中，我们选取彼此之间 曼哈顿距离 最短的工人自行车对 (workeri, bikej) ，
并将其中的自行车分配給工人。

如果有多个 (workeri,bikej) 对之间的 曼哈顿距离 相同，那么我们选择 工人索引最小 的那对。类似地，如果有多种不同的分配方法，
则选择 自行车索引最小 的一对。不断重复这一过程，直到所有工人都分配到自行车为止。

返回长度为 n 的向量 answer，其中 answer[i] 是第 i 位工人分配到的自行车的索引（从 0 开始）。

给定两点 p1 和 p2 之间的 曼哈顿距离 为 Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|。
*/
func assignBikes(workers [][]int, bikes [][]int) []int {

	//n := len(workers)
	//ans := make([]int, n)
	//
	//for i, v := range workers {
	//
	//}
	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(math.Log2(10000000))
}
