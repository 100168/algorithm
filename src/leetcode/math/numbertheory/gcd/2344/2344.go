package main

import (
	"fmt"
	"sort"
)

/*
*
给你两个正整数数组 nums 和 numsDivide 。你可以从 nums 中删除任意数目的元素。

请你返回使 nums 中 最小 元素可以整除 numsDivide 中所有元素的 最少 删除次数。如果无法得到这样的元素，返回 -1 。

如果 y % x == 0 ，那么我们说整数 x 整除 y 。

输入：nums = [2,3,2,4,3], numsDivide = [9,6,9,3,15]
输出：2
解释：
[2,3,2,4,3] 中最小元素是 2 ，它无法整除 numsDivide 中所有元素。
我们从 nums 中删除 2 个大小为 2 的元素，得到 nums = [3,4,3] 。
[3,4,3] 中最小元素为 3 ，它可以整除 numsDivide 中所有元素。
可以证明 2 是最少删除次数。
*/
func minOperations(nums []int, numsDivide []int) int {

	maxGcd := numsDivide[0]
	for _, v := range numsDivide {
		maxGcd = gcd(maxGcd, v)
	}
	sort.Ints(nums)
	for i, v := range nums {
		if maxGcd%v == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func main() {

	fmt.Println(minOperations([]int{2, 3, 2, 4, 3}, []int{9, 6, 9, 3, 15}))
}
