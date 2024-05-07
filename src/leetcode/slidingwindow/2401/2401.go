package main

import "fmt"

/*
给你一个由 正 整数组成的数组 nums 。

如果 nums 的子数组中位于 不同 位置的每对元素按位 与（AND）运算的结果等于 0 ，则称该子数组为 优雅 子数组。

返回 最长 的优雅子数组的长度。

子数组 是数组中的一个 连续 部分。

注意：长度为 1 的子数组始终视作优雅子数组。
*/
func longestNiceSubarray(nums []int) int {

	n := len(nums)

	w := 0

	l := 0
	ans := 0

	for r := 0; r < n; r++ {
		for w&nums[r] != 0 {
			w ^= nums[l]
			l++
		}
		w ^= nums[r]
		ans = max(ans, r-l+1)
	}
	return ans

}

func main() {

	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))
}
