package main

import "fmt"

/*
*
给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。

示例 1:

输入: n = 5
输出: 2
解释: 5 = 2 + 3，共有两组连续整数([5],[2,3])求和后为 5。
示例 2:

输入: n = 9
输出: 3
解释: 9 = 4 + 5 = 2 + 3 + 4
示例 3:

输入: n = 15
输出: 4
解释: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
*/
func consecutiveNumbersSum(n int) int {

	ans := 1
	// (2*s+t-1)*t == 2*n =》n/t = 2*s+t-1 =>n/t-t+1 ==>s = (n/i-i+1)/2
	n *= 2

	for i := 2; i*i <= n; i++ {
		if n%i == 0 && (n/i-i+1)%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(consecutiveNumbersSum(15))
}
