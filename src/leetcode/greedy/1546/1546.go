package main

import "fmt"

/**
给你一个数组 nums 和一个整数 target 。

请你返回 非空不重叠 子数组的最大数目，且每个子数组中数字和都为 target 。



示例 1：

输入：nums = [1,1,1,1,1], target = 2
输出：2
解释：总共有 2 个不重叠子数组（加粗数字表示） [1,1,1,1,1] ，它们的和为目标值 2 。
示例 2：

输入：nums = [-1,3,5,1,4,2,-9], target = 6
输出：2
解释：总共有 3 个子数组和为 6 。
([5,1], [4,2], [3,5,1,4,2,-9]) 但只有前 2 个是不重叠的。
*/

func maxNonOverlapping(nums []int, target int) int {

	ans := 0

	sumMap := make(map[int]int)

	s := 0

	pre := -1
	sumMap[0] = -1

	for i, v := range nums {
		s += v
		if _, ok := sumMap[s-target]; ok && sumMap[s-target] >= pre {
			ans++
			pre = i
		}

		sumMap[s] = i

	}
	return ans

}

func main() {
	//fmt.Println(maxNonOverlapping([]int{1, 1, 1, 1, 1}, 2))
	fmt.Println(maxNonOverlapping([]int{-1, 3, 5, 1, 4, 2, -9}, 6))
}
