package main

import "math"

/*
*
有 n 个城市，按从 0 到 n-1 编号。给你一个边数组 edges，其中
edges[i] = [fromi, toi, weighti] 代表 fromi 和 toi 两个城市之间的双向加权边，距离阈值是一个整数 distanceThreshold。

返回在路径距离限制为 distanceThreshold 以内可到达城市最少的城市。如果有多个这样的城市，则返回编号最大的城市。

注意，连接城市 i 和 j 的路径的距离等于沿该路径的所有边的权重之和。
*/
func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
		f[i][i] = 0
	}
	w := make([][]int, n)
	for i := range w {
		w[i] = make([]int, n)

		for j := range w[i] {
			w[i][j] = math.MaxInt / 2
		}
	}
	for _, v := range edges {
		x, y, d := v[0], v[1], v[2]
		w[x][y] = d
		w[y][x] = d
	}

	f = w
	// i,j k cur = f(i,k,k-1)+f( k,j,k-1),
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = min(f[i][k]+f[k][j], f[i][j])
			}
		}
	}

	minCnt := math.MaxInt

	ans := math.MaxInt
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if f[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt < minCnt || cnt == minCnt && ans > i {
			ans = i
			minCnt = cnt
		}
	}
	return ans
}
