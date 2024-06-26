package main

import (
	"slices"
)

// 给你一个整数数组 nums 和一个 正整数 k 。
//
// 请你统计有多少满足 「 nums 中的 最大 元素」至少出现 k 次的子数组，并返回满足这一条件的子数组的数目。
//
// 子数组是数组中的一个连续元素序列。
//
// 示例 1：
//
// 输入：nums = [1,3,2,3,3], k = 2
// 输出：6
// 解释：包含元素 3 至少 2 次的子数组为：[1,3,2,3]、[1,3,2,3,3]、[3,2,3]、[3,2,3,3]、[2,3,3] 和 [3,3]
// 。
//
// 示例 2：
//
// 输入：nums = [1,4,2,1], k = 3
// 输出：0
// 解释：没有子数组包含元素 4 至少 3 次。
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁶
// 1 <= k <= 10⁵
//
// leetcode submit region begin(Prohibit modification and deletion)
func countSubarrays(nums []int, k int) int64 {

	maxVal := slices.Max(nums)

	cnt := 0
	for _, v := range nums {
		if v == maxVal {
			cnt++
		}
	}

	if cnt < k {
		return 0
	}

	l := 0

	w := 0
	ans := 0
	for _, v := range nums {
		if v == maxVal {
			w++
		}
		for nums[l] != maxVal || w > k {
			if nums[l] == maxVal {
				w--
			}
			l++
		}
		if w == k {
			ans += l + 1
		}
	}
	return int64(ans)
}

func countSubarrays2(nums []int, k int) (ans int64) {
	mx := slices.Max(nums)
	cntMx, left := 0, 0
	for _, x := range nums {
		if x == mx {
			cntMx++
		}
		for cntMx == k {
			if nums[left] == mx {
				cntMx--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func main() {
	println(countSubarrays([]int{1, 3, 2, 3, 3}, 2))
}

//leetcode submit region end(Prohibit modification and deletion)
