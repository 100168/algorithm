package main

import (
	"fmt"
	"math/bits"
)

/*
*
给你一个长度为 3 的整数数组 nums。

现以某种顺序 连接 数组 nums 中所有元素的 二进制表示 ，请你返回可以由这种方法形成的 最大 数值。

注意 任何数字的二进制表示 不含 前导零。

示例 1:

输入: nums = [1,2,3]

输出: 30

解释:

按照顺序 [3, 1, 2] 连接数字的二进制表示，得到结果 "11110"，这是 30 的二进制表示。

示例 2:

输入: nums = [2,8,16]

输出: 1296

解释:

按照顺序 [2, 8, 16] 连接数字的二进制表述，得到结果 "10100010000"，这是 1296 的二进制表示。
*/
func maxGoodNumber(nums []int) int {

	var dfs func(int) int

	dfs = func(i int) int {
		cur := 0
		for j := 0; j < len(nums); j++ {
			if 1<<j&i == 0 {
				l := bits.Len(uint(dfs(1<<j | i)))
				cur = max(cur, nums[j]<<l+dfs(1<<j|i))
			}
		}
		return cur
	}
	return dfs(0)
}

func main() {

	fmt.Println(maxGoodNumber([]int{1, 11, 5}))
}
