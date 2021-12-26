package main

/*给定一个包含n + 1 个整数的数组nums ，其数字都在 1 到 n之间（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-duplicate-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[slow]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		fast = nums[fast]
		slow = nums[slow]

	}

	return slow
}

func swap(a int, b int, nums []int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}
