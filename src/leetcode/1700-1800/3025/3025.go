package main

import (
	"math"
	"sort"
)

/*给你一个  n x 2 的二维数组 points ，它表示二维平面上的一些点坐标，其中 points[i] = [xi, yi] 。

我们定义 x 轴的正方向为 右 （x 轴递增的方向），x 轴的负方向为 左 （x 轴递减的方向）。类似的，我们定义 y 轴的正方向为 上 （y 轴递增的方向），y 轴的负方向为 下 （y 轴递减的方向）。

你需要安排这 n 个人的站位，这 n 个人中包括 Alice 和 Bob 。你需要确保每个点处 恰好 有 一个 人。同时，Alice 想跟 Bob 单独玩耍，所以 Alice 会以 Alice 的坐标为 左上角 ，Bob 的坐标为 右下角 建立一个矩形的围栏（注意，围栏可能 不 包含任何区域，也就是说围栏可能是一条线段）。如果围栏的 内部 或者 边缘 上有任何其他人，Alice 都会难过。

请你在确保 Alice 不会 难过的前提下，返回 Alice 和 Bob 可以选择的 点对 数目。

注意，Alice 建立的围栏必须确保 Alice 的位置是矩形的左上角，Bob 的位置是矩形的右下角。比方说，以 (1, 1) ，(1, 3) ，(3, 1) 和 (3, 3) 为矩形的四个角，给定下图的两个输入，Alice 都不能建立围栏，原因如下：

图一中，Alice 在 (3, 3) 且 Bob 在 (1, 1) ，Alice 的位置不是左上角且 Bob 的位置不是右下角。
图二中，Alice 在 (1, 3) 且 Bob 在 (1, 1) ，Bob 的位置不是在围栏的右下角。*/

func numberOfPairs(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {

		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] > points[j][1]
	})
	ans := 0
	n := len(points)
	for i := range points {
		up := points[i]
		m := math.MinInt
		for j := i + 1; j < n; j++ {
			down := points[j]
			if down[1] <= up[1] && down[1] > m {
				ans++
				m = down[1]
			}
		}
	}
	return ans
}
func main() {
	println(numberOfPairs([][]int{{6, 2}, {4, 4}, {2, 6}}))
}
