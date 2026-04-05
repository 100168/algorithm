package main

func minOperations(nums []int, k int) int {

	n := len(nums)

	target := int64(0)
	for i := 1; i <= k; i++ {
		target |= 1 << i
	}
	ans := int64(0)
	for i := n - 1; i >= 0; i-- {
		ans |= 1 << nums[i]
		if ans&target == target {
			return n - i
		}

	}
	return 0
}
