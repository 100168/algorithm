package main

func countMatchingSubarrays(nums []int, pattern []int) int {

	n := len(nums)
	m := len(pattern)

	ans := 0
next:
	for i := m; i < n; i++ {

		first := i - (m)
		for j, k := first+1, 0; j <= i; j++ {
			if pattern[k] == 0 && nums[j] != nums[j-1] {
				continue next
			}
			if pattern[k] == 1 && nums[j] <= nums[j-1] {
				continue next
			}
			if pattern[k] == -1 && nums[j] >= nums[j-1] {
				continue next
			}
			k++
		}
		ans++
	}

	return ans
}
