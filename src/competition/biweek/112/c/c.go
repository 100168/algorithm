package main

//6989. 几乎唯一子数组的最大和 显示英文描述
//通过的用户数0
//尝试过的用户数0
//用户总通过次数0
//用户总提交次数0
//题目难度Medium
//给你一个整数数组 nums 和两个正整数 m 和 k 。
//
//请你返回 nums 中长度为 k 的 几乎唯一 子数组的 最大和 ，如果不存在几乎唯一子数组，请你返回 0 。
//
//如果 nums 的一个子数组有至少 m 个互不相同的元素，我们称它是 几乎唯一 子数组。
//
//子数组指的是一个数组中一段连续 非空 的元素序列。

func maxSum(nums []int, m int, k int) int64 {

	windSum := int64(0)
	l := 0
	windNum := make(map[int]int)
	ans := int64(0)
	for i, num := range nums {
		windSum += int64(num)
		windNum[num]++
		if i-l+1 > k {
			windSum -= int64(nums[l])
			windNum[nums[l]]--
			if windNum[nums[l]] == 0 {
				delete(windNum, nums[l])
			}
			l++
		}

		if len(windNum) >= m {
			ans = max(ans, windSum)
		}

	}

	return ans
}
