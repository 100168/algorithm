package main

import "fmt"

/*现有一个按 升序 排列的整数数组 nums ，其中每个数字都 互不相同 。

给你一个整数 k ，请你找出并返回从数组最左边开始的第 k 个缺失数字。

*/

/**到nums[idx]一共缺失的数字*/
func missing(idx int, nums []int) int {

	return nums[idx] - nums[0] - idx
}

// 线性遍历
func missingE(nums []int, k int) int {
	l := len(nums) - 1
	if k > missing(l, nums) {
		return nums[l] + k - missing(l, nums)
	}

	idx := 1
	for ; missing(idx, nums) < k; idx++ {

	}
	return nums[idx-1] + k - missing(idx-1, nums)
}

// 二分搜索
func missingElement(nums []int, k int) int {

	l := len(nums) - 1

	if k > missing(l, nums) {
		return nums[l] + k - missing(l, nums)
	}

	left := 0
	right := l

	middle := left + (right - left>>1)

	for left < right {
		if k > missing(middle, nums) {
			left = middle + 1
		} else {
			right = middle
		}
		middle = left + (right - left>>1)
	}

	return nums[right-1] + k - missing(right-1, nums)
}

func main() {

	nums := []int{1, 2, 4}
	element := missingElement(nums, 3)

	fmt.Println(element)
}
