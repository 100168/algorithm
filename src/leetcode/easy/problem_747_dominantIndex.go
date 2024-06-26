package main

import (
	"fmt"
)

/*
*给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。

请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func dominantIndex(nums []int) int {
	m1, m2, idx := -1, -1, 0
	for i, num := range nums {
		if num > m1 {
			m1, m2, idx = num, m1, i
		} else if num > m2 {
			m2 = num
		}
	}
	if m1 >= m2*2 {
		return idx
	}
	return -1
}

func main() {
	nums := []int{0, 1}

	index := dominantIndex(nums)

	fmt.Println(index)
}
