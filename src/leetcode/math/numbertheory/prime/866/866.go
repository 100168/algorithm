package main

import (
	"fmt"
	"math"
)

/*
*
给你一个整数 n ，返回大于或等于 n 的最小 回文质数。
一个整数如果恰好有两个除数：1 和它本身，那么它是 质数 。注意，1 不是质数。

例如，2、3、5、7、11 和 13 都是质数。
一个整数如果从左向右读和从右向左读是相同的，那么它是 回文数 。

例如，101 和 12321 都是回文数。
测试用例保证答案总是存在，并且在 [2, 2 * 108] 范围内。

输入：n = 6
输出：7
示例 2：

输入：n = 8
输出：11
示例 3：

输入：n = 13
输出：101
*/
func primePalindrome(n int) int {

	isPrime := func(x int) bool {

		if x < 2 {
			return false
		}
		for i := 2; i*i <= x; i++ {

			if x%i == 0 {
				return false
			}

		}
		return true
	}

	for i := 1; i < 9; i++ {
		up := int(math.Pow10((i + 1) / 2))
		for j := up / 10; j < up; j++ {
			x := j
			c := j
			if i%2 != 0 {
				x /= 10
			}

			for x > 0 {
				c = c*10 + x%10
				x /= 10
			}

			if c >= n && isPrime(c) {
				return c
			}

		}
	}

	return -1

}
func main() {
	fmt.Println(primePalindrome(80283903))
	fmt.Println(primePalindrome(13))
	fmt.Println(primePalindrome(6))
	fmt.Println(primePalindrome(8))
	fmt.Println(primePalindrome(1))
}
