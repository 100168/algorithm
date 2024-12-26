package main

/*
*
给你一个仅由数字 0 - 9 组成的字符串 num。如果偶数下标处的数字之和等于奇数下标处的数字之和，则认为该数字字符串是一个 平衡字符串。

如果 num 是一个 平衡字符串，则返回 true；否则，返回 false。

示例 1：

输入：num = "1234"

输出：false

解释：

偶数下标处的数字之和为 1 + 3 = 4，奇数下标处的数字之和为 2 + 4 = 6。
由于 4 不等于 6，num 不是平衡字符串。
*/
func isBalanced(num string) bool {

	s := 0
	for i, v := range num {

		c := int(v - '0')

		if i%2 == 0 {
			s += c
		} else {
			s -= c
		}
	}
	return s == 0
}
