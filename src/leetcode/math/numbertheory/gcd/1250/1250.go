package main

/*
*
给你一个正整数数组 nums，你需要从中任选一些子集，然后将子集中每一个数乘以一个 任意整数，并求出他们的和。

假如该和结果为 1，那么原数组就是一个「好数组」，则返回 True；否则请返回 False。

示例 1：

输入：nums = [12,5,7,23]
输出：true
解释：挑选数字 5 和 7。
5*3 + 7*(-2) = 1
示例 2：

输入：nums = [29,6,10]
输出：true
解释：挑选数字 29, 6 和 10。
29*1 + 6*(-3) + 10*(-1) = 1
示例 3：

输入：nums = [3,6]
输出：false
*/
func isGoodArray(nums []int) bool {

	//nums[i]
	//29*a+6*b = 1 ===> b = (1-29*a)/6

	preGcd := nums[0]
	for i := 0; i < len(nums); i++ {
		preGcd = gcd(preGcd, nums[i])
		if preGcd == 1 {
			return true
		}
	}
	return false
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
