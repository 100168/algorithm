package main

import (
	"fmt"
	"sort"
)

/*
*

给你一个二维数组 points 和一个字符串 s ，其中 points[i] 表示第 i 个点的坐标，s[i] 表示第 i 个点的 标签 。

如果一个正方形的中心在 (0, 0) ，所有边都平行于坐标轴，且正方形内 不 存在标签相同的两个点，那么我们称这个正方形是 合法 的。

请你返回 合法 正方形中可以包含的 最多 点数。

注意：

如果一个点位于正方形的边上或者在边以内，则认为该点位于正方形内。
正方形的边长可以为零。
*/
func maxPointsInsideSquare(points [][]int, s string) int {

	type pair struct {
		x int
		y uint8
	}

	nums := make([]pair, 0)

	for i, p := range points {
		curMax := max(abs(p[0]), abs(p[1]))
		nums = append(nums, pair{curMax, s[i]})
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].x < nums[j].x
	})

	cntMap := make(map[uint8]bool)

	ans := 0

	for i := 0; i < len(nums); i++ {

		v := nums[i]
		ans = len(cntMap)
		j := i
		for ; j < len(nums); j++ {
			if nums[j].x != v.x {
				break
			}
			if !cntMap[nums[j].y] {
				cntMap[nums[j].y] = true
			} else {
				return ans
			}
		}
		i = j - 1
	}
	return len(cntMap)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(maxPointsInsideSquare([][]int{{-35, -3}, {17, 28}, {28, -28}, {25, -1}, {25, -16}, {1, -21}}, "ffcbea"))
}
