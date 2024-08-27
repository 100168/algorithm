package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
给你一个正整数数组 nums 。

如果我们执行以下操作 至多一次 可以让两个整数 x 和 y 相等，那么我们称这个数对是 近似相等 的：

选择 x 或者 y  之一，将这个数字中的两个数位交换。
请你返回 nums 中，下标 i 和 j 满足 i < j 且 nums[i] 和 nums[j] 近似相等 的数对数目。

注意 ，执行操作后一个整数可以有前导 0 。



示例 1：

输入：nums = [3,12,30,17,21]

输出：2

解释：

近似相等数对包括：

3 和 30 。交换 30 中的数位 3 和 0 ，得到 3 。
12 和 21 。交换12 中的数位 1 和 2 ，得到 21 。
*/

func countPairs(nums []int) int {

	sort.Ints(nums)
	ans := 0

	for i, v1 := range nums {

	next:
		for _, v2 := range nums[i+1:] {

			if v2 == v1 {
				ans++
				continue next
			}
			itoa := strconv.Itoa(v2)

			bytes := []byte(itoa)

			for c := range bytes {
				for cc := c + 1; cc < len(bytes); cc++ {
					bytes[c], bytes[cc] = bytes[cc], bytes[c]
					s := string(bytes)
					atoi, _ := strconv.Atoi(s)

					if atoi == v1 {
						ans++
						continue next
					}
					bytes[c], bytes[cc] = bytes[cc], bytes[c]
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countPairs([]int{5, 12, 8, 5, 5, 1, 20, 3, 10, 10, 5, 5, 5, 5, 1}))
}
