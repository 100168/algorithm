package main

/**
给你两个长度为 n 的整数数组 nums 和 forbidden。

你可以执行以下操作任意次（包括零次）：

选择两个不同下标 i 和 j，然后交换 nums[i] 和 nums[j]。
返回使得对于每个下标 i，nums[i] 不等于 forbidden[i] 所需的最小交换次数。如果无论如何都无法满足条件，返回 -1。

示例 1：

输入： nums = [1,2,3], forbidden = [3,2,1]

输出： 1

解释：

一种最优的交换方案：

选择 nums 中下标 i = 0 和 j = 1，交换它们，得到 nums = [2, 1, 3]。
交换完成后，对于每个下标 i，nums[i] 都不等于 forbidden[i]。
示例 2：

输入： nums = [4,6,6,5], forbidden = [4,6,5,5]

输出： 2

解释：

一种最优的交换方案：

选择 nums 中下标 i = 0 和 j = 2，交换它们，得到 nums = [6, 6, 4, 5]。
选择 nums 中下标 i = 1 和 j = 3，交换它们，得到 nums = [6, 5, 4, 6]。
交换完成后，对于每个下标 i，nums[i] 都不等于 forbidden[i]。
示例 3：

输入： nums = [7,7], forbidden = [8,7]

输出： -1

解释：

无法通过任何交换使得 nums[i] 对于所有下标 i 都不等于 forbidden[i]。

示例 4：

输入： nums = [1,2], forbidden = [2,1]

输出： 0

解释：

无需交换，因为对于每个下标 i，nums[i] 已经不等于 forbidden[i]，因此答案是 0。



提示：

1 <= n == nums.length == forbidden.length <= 10^5
1 <= nums[i], forbidden[i] <= 10^9
*/

func minSwaps(nums []int, forbidden []int) int {
	n := len(nums)
	numCnt := map[int]int{}
	forbCnt := map[int]int{}
	for i := 0; i < n; i++ {
		numCnt[nums[i]]++
		forbCnt[forbidden[i]]++
	}

	for v, c := range numCnt {
		if c > n-forbCnt[v] {
			return -1
		}
	}
	total := 0
	cnt := map[int]int{}
	maxC := 0
	for i := 0; i < n; i++ {
		if nums[i] == forbidden[i] {
			v := nums[i]
			total++
			cnt[v]++
			if cnt[v] > maxC {
				maxC = cnt[v]
			}
		}
	}

	if total == 0 {
		return 0
	}

	res := (total + 1) / 2
	if maxC > res {
		res = maxC
	}

	return res
}

func main() {
	println(minSwaps([]int{4, 6, 6, 5}, []int{4, 6, 5, 5}))
	println(minSwaps([]int{7, 7}, []int{8, 7}))
	println(minSwaps([]int{1, 2}, []int{2, 1}))
}
