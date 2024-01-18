//There is an integer array nums sorted in ascending order (with distinct
//values).
//
// Prior to being passed to your function, nums is possibly rotated at an
//unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k]
//, nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For
//example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0
//,1,2].
//
// Given the array nums after the possible rotation and an integer target,
//return the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
//
// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1
//
// Example 3:
// Input: nums = [1], target = 0
//Output: -1
//
//
// Constraints:
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -10⁴ <= target <= 10⁴
//
//
// Related Topics Array Binary Search 👍 25245 👎 1495

package _3

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1
	n := len(nums) - 1

	// Input: nums = [4,5,6,7,0,1,2], target = 3
	//这里用<=是因为每次查找都不包括中间的数
	for l <= r {
		m := (r + l) / 2
		if nums[m] == target {
			return m
		}

		//1.如果中间数大于等于最左边的数，说明左半区间是递增的
		//2.否则说明右半区间是递增的
		if nums[m] >= nums[0] {
			//1.目标数比当前数小，并且大于等于最左边数，往左半区间找
			//2.否则说明当前数在右半区间
			if nums[0] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			//1.到这说明右边是递增的
			//2.如果目标数大于当前数，并且小于等于最右边数，说明要往右半区间找
			//3.否则往左半区间找
			if target > nums[m] && target <= nums[n] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	//说明没找到，返回-1

	return -1

}

//leetcode submit region end(Prohibit modification and deletion)
