package main

import "fmt"

/*
*
给你一个长度为 n 下标从 0 开始的整数数组 maxHeights 。

你的任务是在坐标轴上建 n 座塔。第 i 座塔的下标为 i ，高度为 heights[i] 。

如果以下条件满足，我们称这些塔是 美丽 的：

1 <= heights[i] <= maxHeights[i]
heights 是一个 山脉 数组。
如果存在下标 i 满足以下条件，那么我们称数组 heights 是一个 山脉 数组：

对于所有 0 < j <= i ，都有 heights[j - 1] <= heights[j]
对于所有 i <= k < n - 1 ，都有 heights[k + 1] <= heights[k]
请你返回满足 美丽塔 要求的方案中，高度和的最大值 。

输入：maxHeights = [5,3,4,1,1]
输出：13
解释：和最大的美丽塔方案为 heights = [5,3,3,1,1] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]
- heights 是个山脉数组，峰值在 i = 0 处。
13 是所有美丽塔方案中的最大高度和。

5,3,
*/
func maximumSumOfHeights(maxHeights []int) int64 {

	n := len(maxHeights)
	sumL := make([]int, n)
	sumR := make([]int, n)
	st := make([]int, 0)
	s := 0
	for i := 0; i < n; i++ {
		for len(st) > 0 && maxHeights[st[len(st)-1]] > maxHeights[i] {
			pop := st[len(st)-1]
			st = st[:len(st)-1]
			l := -1
			if len(st) > 0 {
				l = st[len(st)-1]
			}
			s -= maxHeights[pop] * (pop - l)
		}

		l := -1

		if len(st) > 0 {
			l = st[len(st)-1]
		}
		st = append(st, i)
		s += maxHeights[i] * (i - l)
		sumL[i] = s
	}

	st = st[:0]
	s = 0
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && maxHeights[st[len(st)-1]] > maxHeights[i] {
			pop := st[len(st)-1]
			st = st[:len(st)-1]
			l := n
			if len(st) > 0 {
				l = st[len(st)-1]
			}
			s -= maxHeights[pop] * (l - pop)
		}

		l := n

		if len(st) > 0 {
			l = st[len(st)-1]
		}
		st = append(st, i)
		s += maxHeights[i] * (l - i)
		sumR[i] = s
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, sumL[i]+sumR[i]-maxHeights[i])
	}
	return int64(ans)

}

func main() {
	fmt.Println(maximumSumOfHeights([]int{5, 3, 4, 1, 1}))
}
