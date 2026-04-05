package main

/*
给你一个长度为 n 下标从 0 开始的整数数组 maxHeights 。

你的任务是在坐标轴上建 n 座塔。第 i 座塔的下标为 i ，高度为 heights[i] 。

如果以下条件满足，我们称这些塔是 美丽 的：

1 <= heights[i] <= maxHeights[i]
heights 是一个 山状 数组。
如果存在下标 i 满足以下条件，那么我们称数组 heights 是一个 山状 数组：

对于所有 0 < j <= i ，都有 heights[j - 1] <= heights[j]
对于所有 i <= k < n - 1 ，都有 heights[k + 1] <= heights[k]
请你返回满足 美丽塔 要求的方案中，高度和的最大值
*/
func maximumSumOfHeights(maxHeights []int) int64 {

	n := len(maxHeights)
	ans := int64(0)
	for i := 0; i < n; i++ {

		cur := int64(maxHeights[i])
		leftMinValue := cur
		rightMinValue := cur

		for j := i - 1; j >= 0; j-- {
			if int64(maxHeights[j]) < leftMinValue {
				leftMinValue = int64(maxHeights[j])
			}
			cur += leftMinValue

		}

		for j := i + 1; j < n; j++ {
			if int64(maxHeights[j]) < rightMinValue {
				rightMinValue = int64(maxHeights[j])
			}
			cur += rightMinValue
		}

		if ans < cur {
			ans = cur
		}
	}
	return ans
}
