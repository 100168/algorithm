package main

import "fmt"

/*
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//单调队列，加滑动窗口
func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := 0
	n := len(nums)

	var stack []int
	var ans []int
	for right < n {
		if k > 0 {
			for len(stack) > 0 && nums[right] > stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, nums[right])
			right++
			k--
		}
		if k == 0 {
			ans = append(ans, stack[0])
			if nums[left] == stack[0] {
				stack = stack[1:]
			}
			left++
			k++
		}
	}

	return ans
}

func maxSlidingWindow2(nums []int, k int) []int {
	var q []int
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7, 23, 3, 4, 3, 234, 24, 453, 3, 3, 23, 4, 45, 45, 56, 5, 4, 454, 65, 7}

	window := maxSlidingWindow(a, 3)
	fmt.Println(window)

	window = maxSlidingWindow2(a, 3)
	fmt.Println(window)
}
