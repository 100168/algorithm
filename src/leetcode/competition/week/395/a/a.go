package main

// 给你两个长度相等的数组 nums1 和 nums2。
//
// 数组 nums1 中的每个元素都与变量 x 所表示的整数相加。如果 x 为负数，则表现为元素值的减少。
//
// 在与 x 相加后，nums1 和 nums2 相等 。当两个数组中包含相同的整数，并且这些整数出现的频次相同时，两个数组 相等 。
//
// 返回整数 x 。
func addedInteger(nums1 []int, nums2 []int) int {

	n1 := len(nums1)
	s1 := 0
	s2 := 0
	for _, v := range nums1 {
		s1 += v
	}

	for _, v := range nums2 {
		s2 += v
	}

	return (s2 - s1) / n1

}
