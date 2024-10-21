package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个整数数组 nums，和一个整数 k 。

对于每个下标 i（0 <= i < nums.length），将 nums[i] 变成 nums[i] + k 或 nums[i] - k 。

nums 的 分数 是 nums 中最大元素和最小元素的差值。

在更改每个下标对应的值之后，返回 nums 的最小 分数 。

示例 1：

输入：nums = [1], k = 0
输出：0
解释：分数 = max(nums) - min(nums) = 1 - 1 = 0 。

示例 2：

输入：nums = [0,10], k = 2
输出：6
解释：将数组变为 [2, 8] 。分数 = max(nums) - min(nums) = 8 - 2 = 6 。
示例 3：

输入：nums = [1,3,6], k = 3
输出：3
解释：将数组变为 [4, 6, 3] 。分数 = max(nums) - min(nums) = 6 - 3 = 3 。
*/
func smallestRangeII(nums []int, k int) int {

	//1.全+
	//2.减
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i, a := range nums[:n-1] {
		b := nums[i+1]
		high := max(nums[n-1]-k, a+k)
		low := min(b-k, nums[0]+k)
		ans = min(ans, high-low)

	}

	return ans
}

func smallestRangeII2(nums []int, k int) int {

	//1.全+
	//2.减
	for i := range nums {
		nums[i] += k
	}
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]

	for i := n - 1; i > 0; i-- {
		low := min(nums[i]-2*k, nums[0])
		high := max(nums[n-1]-2*k, nums[i-1])
		ans = min(ans, high-low)
	}

	return ans
}

func main() {
	//fmt.Println(smallestRangeII([]int{0, 10}, 2))
	//fmt.Println(smallestRangeII([]int{1, 3, 6}, 3))
	fmt.Println(smallestRangeII([]int{1, 4, 6, 4}, 3))
}
