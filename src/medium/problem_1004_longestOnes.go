package main

/*给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。*/

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0

	flag := 0

	ans := 0
	for right < len(nums) {

		if flag < k+1 {
			if nums[right] == 0 {
				flag++
			}
			right++
		}

		if flag == k+1 {
			if nums[left] == 0 {
				flag--
			}
			left++
		}

		if ans < right-left {
			ans = right - left
		}

	}

	return ans
}
