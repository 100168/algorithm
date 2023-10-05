package main

import (
	"sort"
)

/*
给你一个二维整数数组 circles ，其中 circles[i] = [xi, yi, ri] 表示网格上圆心为 (xi, yi) 且半径为 ri 的第 i 个圆，返回出现在 至少一个 圆内的 格点数目 。

注意：
格点 是指整数坐标对应的点。
圆周上的点 也被视为出现在圆内的点。
https://leetcode.cn/problems/count-lattice-points-inside-a-circle/description/
*/
func countLatticePoints(circles [][]int) int {

	ans := 0

	sort.Slice(circles, func(i, j int) bool {
		return circles[i][2] > circles[j][2]
	})
	for i := 0; i <= 200; i++ {
		for j := 0; j <= 200; j++ {
			for _, v := range circles {
				if (i-v[0])*(i-v[0])+(j-v[1])*(j-v[1]) <= v[2]*v[2] {
					ans++
					break
				}
			}
		}
	}
	return ans
}
