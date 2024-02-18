//There is an integer array nums sorted in non-decreasing order (not
//necessarily with distinct values).
//
// Before being passed to your function, nums is rotated at an unknown pivot
//index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1]
//, ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0
//,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,
//2,4,4].
//
// Given the array nums after the rotation and an integer target, return true
//if target is in nums, or false if it is not in nums.
//
// You must decrease the overall operation steps as much as possible.
//
//
// Example 1:
// Input: nums = [2,5,6,0,0,1,2], target = 0
//Output: true
//
// Example 2:
// Input: nums = [2,5,6,0,0,1,2], target = 3
//Output: false
//
//
// Constraints:
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums is guaranteed to be rotated at some pivot.
// -10⁴ <= target <= 10⁴
//
//
//
// Follow up: This problem is similar to Search in Rotated Sorted Array, but
//nums may contain duplicates. Would this affect the runtime complexity? How and why?
//
//
// Related Topics Array Binary Search 👍 8173 👎 1003

package main

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) bool {

	n := len(nums)
	l := 0
	r := n - 1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		}
		//相等说明有重复数据，一直往右边找
		if nums[l] == nums[r] {
			l++
			continue
		}
		if nums[l] <= nums[m] {

			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[n-1] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
