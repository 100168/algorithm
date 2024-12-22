package a

import "sort"

/**
给你一个下标从 0 开始长度为 3 的整数数组 nums ，需要用它们来构造三角形。

如果一个三角形的所有边长度相等，那么这个三角形称为 equilateral 。
如果一个三角形恰好有两条边长度相等，那么这个三角形称为 isosceles 。
如果一个三角形三条边的长度互不相同，那么这个三角形称为 scalene 。
如果这个数组无法构成一个三角形，请你返回字符串 "none" ，否则返回一个字符串表示这个三角形的类型。



示例 1：

输入：nums = [3,3,3]
输出："equilateral"
解释：由于三条边长度相等，所以可以构成一个等边三角形，返回 "equilateral" 。

*/

func triangleType(nums []int) string {

	sort.Ints(nums)

	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}

	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	}

	if nums[0] == nums[1] || nums[1] == nums[2] {

		return "isosceles"
	}

	return "scalene"
}
