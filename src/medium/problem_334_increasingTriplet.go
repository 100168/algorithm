package main

import "math"

/**给你一个整数数组nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k)且满足 i < j < k ，使得nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-triplet-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func increasingTriplet(nums []int) bool {

	n := len(nums)

	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	rightMax := make([]int, n)

	leftMin[0] = nums[0]
	rightMax[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if nums[i] < leftMin[i-1] {
			leftMin[i] = nums[i]
		} else {
			leftMin[i] = leftMin[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		if nums[i] > rightMax[i+1] {
			rightMax[i] = nums[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	for i := 1; i < n; i++ {
		if leftMin[i] < nums[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false

}

//贪心算法，
func increasingTriplet2(nums []int) bool {
	n := len(nums)

	first := nums[0]
	second := math.MaxInt64

	for i := 1; i < n; i++ {
		num := nums[i]

		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}

	}

	return false
}
