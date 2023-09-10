package main

import (
	"strconv"
	"strings"
)

/*
给你一个下标从 0 开始的字符串 num ，表示一个非负整数。

在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。

返回最少需要多少次操作可以使 num 变成特殊数字。

如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。
*/
func minimumOperations(num string) int {
	n := len(num)
	str := []string{"25", "50", "75", "00"}
	ans := n
	for _, s := range str {
		ans = min(getAns(s, num, n), ans)
	}
	return min(ans, n-strings.Count(num, strconv.Itoa(0)))
}

func getAns(di string, num string, n int) int {

	i := strings.LastIndexByte(num, di[1])
	if i == -1 {
		return n
	}
	i = strings.LastIndexByte(num[0:i], di[0])
	if i == -1 {
		return n
	}
	return n - i - 2
}
func main() {
	println(minimumOperations("50"))
}
