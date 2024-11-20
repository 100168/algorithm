package main

/*
*
给你两个整数 n 和 t 。请你返回大于等于 n 的 最小 整数，且该整数的 各数位之积 能被 t 整除。

示例 1：

输入：n = 10, t = 2

输出：10

解释：

10 的数位乘积为 0 ，可以被 2 整除，所以它是大于等于 10 且满足题目要求的最小整数。

示例 2：

输入：n = 15, t = 3

输出：16

解释：

16 的数位乘积为 6 ，可以被 3 整除，所以它是大于等于 15 且满足题目要求的最小整数。
*/
func smallestNumber(n int, t int) int {

	for {
		pre := n

		cur := 1
		for pre > 0 {
			cur *= pre % 10
			pre /= 10
		}
		if cur%t == 0 {
			return n
		}
		n++
	}
}
