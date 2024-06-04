package main

import "fmt"

/*
*
给你一个二元数组 nums 。

如果数组中的某个子数组 恰好 只存在 一 个值为 1 的元素，则认为该子数组是一个 好子数组 。

请你统计将数组 nums 划分成若干 好子数组 的方法数，并以整数形式返回。
由于数字可能很大，返回其对 109 + 7 取余 之后的结果。

子数组是数组中的一个连续 非空 元素序列。
*/
func numberOfGoodSubarraySplits(nums []int) int {
	mod := int(1e9 + 7)

	index := make([]int, 0)

	for i := range nums {

		if nums[i] == 1 {
			index = append(index, i)
		}
	}

	if len(index) <= 1 {
		return len(index)
	}
	ans := 1
	pre := -1
	for _, v := range index {
		if pre >= 0 {
			ans = ans * (v - pre) % mod
		}
		pre = v

	}
	return ans % mod

}

func numberOfGoodSubarraySplits2(nums []int) int {
	const mod int = 1e9 + 7
	ans, pre := 1, -1
	for i, x := range nums {
		if x > 0 {
			if pre >= 0 {
				ans = ans * (i - pre) % mod
			}
			pre = i
		}
	}
	if pre < 0 { // 整个数组都是 0，没有好子数组
		return 0
	}
	return ans
}

func main() {
	fmt.Println(numberOfGoodSubarraySplits([]int{1, 0, 0, 0, 0, 0, 1, 0, 1}))
	fmt.Println(numberOfGoodSubarraySplits([]int{0, 0, 1, 1}))
}
