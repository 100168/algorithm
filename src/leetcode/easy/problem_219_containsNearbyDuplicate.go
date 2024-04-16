package main

func containsNearbyDuplicate2(nums []int, k int) bool {

	n := len(nums)

	numsMap := make(map[int]int)

	for i := 0; i < n; i++ {

		num := numsMap[nums[i]]
		if num != 0 {
			if abs(num, i+1) <= k {
				return true
			}
		}
		numsMap[nums[i]] = i + 1
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {

	n := len(nums)

	if k == 0 {
		return false
	}
	left := 0
	numSet := make(map[int]bool)

	for i := 0; i < n; i++ {
		if numSet[nums[i]] {
			return true
		}
		if i-left == k {
			if len(numSet) > 0 {
				delete(numSet, nums[left])
				left++
			}
		}
		numSet[nums[i]] = true
	}

	return false
}
func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
