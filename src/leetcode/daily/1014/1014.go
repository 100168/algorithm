package main

import "fmt"

/**
给你一个正整数数组 values，其中 values[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的 距离 为 j - i。

一对景点（i < j）组成的观光组合的得分为 values[i] + values[j] + i - j ，也就是景点的评分之和 减去 它们两者之间的距离。

返回一对观光景点能取得的最高分。



示例 1：

输入：values = [8,1,5,2,6]
输出：11
解释：i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11
示例 2：

输入：values = [1,2]
输出：2


values[i1]-values[i2] + i1 -i2  > values[j2]-values[j1] - j2+j1


i-j

values[i]+values[j]

*/

func maxScoreSightseeingPair(values []int) int {

	ans := 0

	l := 0
	for i, v := range values[1:] {
		i++
		ans = max(ans, v+values[l]+l-i)
		if v >= values[l]+(l-i) {
			l = i
		}
	}
	return ans
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{7, 1, 10, 6, 3, 5, 10, 2}))
}
