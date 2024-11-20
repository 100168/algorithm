package main

import "fmt"

/*
*
给你三个 正 整数 num1 ，num2 和 num3 。

数字 num1 ，num2 和 num3 的数字答案 key 是一个四位数，定义如下：

一开始，如果有数字 少于 四位数，给它补 前导 0 。
答案 key 的第 i 个数位（1 <= i <= 4）为 num1 ，num2 和 num3 第 i 个数位中的 最小 值。
请你返回三个数字 没有 前导 0 的数字答案。

示例 1：

输入：num1 = 1, num2 = 10, num3 = 1000

输出：0

解释：

补前导 0 后，num1 变为 "0001" ，num2 变为 "0010" ，num3 保持不变，为 "1000" 。

数字答案 key 的第 1 个数位为 min(0, 0, 1) 。
数字答案 key 的第 2 个数位为 min(0, 0, 0) 。
数字答案 key 的第 3 个数位为 min(0, 1, 0) 。
数字答案 key 的第 4 个数位为 min(1, 0, 0) 。
所以数字答案为 "0000" ，也就是 0 。
*/

func generateKey(num1 int, num2 int, num3 int) int {

	ans := 0

	base := 1

	for i := 0; i < 4; i++ {
		cur := min(num1%10, num2%10, num3%10)
		num1, num2, num3 = num1/10, num2/10, num3/10
		ans += base * cur
		base *= 10
	}
	return ans
}
func main() {

	fmt.Println(generateKey(987, 879, 798))
}
