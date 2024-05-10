package main

/*
给你一个下标从 0 开始的 正 整数数组 nums 。

如果 nums 的一个子数组满足：移除这个子数组后剩余元素 严格递增 ，那么我们称这个子数组为 移除递增 子数组。
比方说，[5, 3, 4, 6, 7] 中的 [3, 4] 是一个移除递增子数组，因为移除该子数组后，[5, 3, 4, 6, 7] 变为 [5, 6, 7] ，是严格递增的。

请你返回 nums 中 移除递增 子数组的总数目。

注意 ，剩余元素为空的数组也视为是递增的。

子数组 指的是一个数组中一段连续的元素序列。*/

func incremovableSubarrayCount(nums []int) int64 {

	n := len(nums)

	mostLeft := 0
	mostRight := n - 1

	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			break
		}
		mostLeft = i
	}

	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			break
		}
		mostRight = i
	}

	if mostLeft == n-1 {
		return int64(n * (n + 1) / 2)
	}
	ans := mostLeft + 2
	for r := n - 1; r >= mostRight; r-- {
		for mostLeft >= 0 && nums[mostLeft] >= nums[r] {
			mostLeft--
		}
		ans += mostLeft + 2
	}
	return int64(ans)
}

func main() {
	println(incremovableSubarrayCount([]int{6, 5, 7, 8}))
}
