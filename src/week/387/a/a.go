package main

/*
3069. 将元素分配到两个数组中 I 显示英文描述
通过的用户数2929
尝试过的用户数2967
用户总通过次数2957
用户总提交次数3439
题目难度Easy
给你一个下标从 1 开始、包含 不同 整数的数组 nums ，数组长度为 n 。

你需要通过 n 次操作，将 nums 中的所有元素分配到两个数组 arr1 和 arr2 中。在第一次操作中，将 nums[1] 追加到 arr1 。在第二次操作中，将 nums[2] 追加到 arr2 。之后，在第 i 次操作中：

如果 arr1 的最后一个元素 大于 arr2 的最后一个元素，就将 nums[i] 追加到 arr1 。否则，将 nums[i] 追加到 arr2 。
通过连接数组 arr1 和 arr2 形成数组 result 。例如，如果 arr1 == [1,2,3] 且 arr2 == [4,5,6] ，那么 result = [1,2,3,4,5,6] 。

返回数组 result 。*/

func resultArray(nums []int) []int {

	n := len(nums)
	a := make([]int, 0)
	b := make([]int, 0)
	a = append(a, nums[0])
	b = append(b, nums[1])
	for i := 2; i < n; i++ {
		if a[len(a)-1] > b[len(b)-1] {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}
	ans := make([]int, n)
	index := 0
	for _, v := range a {
		ans[index] = v
		index++
	}
	for _, v := range b {
		ans[index] = v
		index++
	}
	return ans
}
