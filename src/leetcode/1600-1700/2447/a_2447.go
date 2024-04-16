package main

/*
已解答
中等
相关标签
相关企业
提示
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 nums 的子数组中元素的最大公因数等于 k 的子数组数目。

子数组 是数组中一个连续的非空序列。

数组的最大公因数 是能整除数组中所有元素的最大整数。

示例 1：

输入：nums = [9,3,1,2,6,3], k = 3
输出：4
解释：nums 的子数组中，以 3 作为最大公因数的子数组如下：
- [9,3,1,2,6,3]
- [9,3,1,2,6,3]
- [9,3,1,2,6,3]
- [9,3,1,2,6,3]
示例 2：

输入：nums = [4], k = 7
输出：0
解释：不存在以 7 作为最大公因数的子数组。
*/
func subarrayGCD(nums []int, k int) int {

	n := len(nums)
	gcdList := make([]int, 0)
	ans := 0

	for i := 0; i < n; i++ {
		gcdNewList := make([]int, 0)
		gcdList = append(gcdList, nums[i])
		for j := 0; j < len(gcdList); j++ {
			curGcd := gcd(gcdList[j], nums[i])
			if curGcd%k != 0 {
				continue
			}
			if curGcd < k {
				continue
			}
			if curGcd >= k {
				gcdNewList = append(gcdNewList, curGcd)
			}
			if curGcd == k {
				ans++
			}
		}
		gcdList = gcdNewList

	}

	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
