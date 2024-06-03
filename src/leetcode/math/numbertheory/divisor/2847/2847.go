package main

import (
	"strconv"
)

/*
*
给定一个 正 整数 n，返回一个字符串，表示 最小的正整数，使其各位数字的乘积等于 n ，
如果不存在这样的数字，则返回 "-1" 。

跟625一样的题目
*/
func smallestNumber(num int64) string {
	ans := ""

	if num < 10 {
		return strconv.Itoa(int(num))
	}

	for i := int64(9); i > 1; i-- {
		for num%i == 0 {
			num /= i
			ans = strconv.Itoa(int(i)) + ans
		}
	}

	if num > 1 {
		return "-1"
	}
	return ans
}
