package main

import (
	"fmt"
	"strings"
)

/**
给你一个下标从 0 开始的字符串 num ，表示一个非负整数。

在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。

返回最少需要多少次操作可以使 num 变成特殊数字。

如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。





示例 1：

输入：num = "2245047"
输出：2
解释：删除数字 num[5] 和 num[6] ，得到数字 "22450" ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 2 位数字。
示例 2：

输入：num = "2908305"
输出：3
解释：删除 num[3]、num[4] 和 num[6] ，得到数字 "2900" ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 3 位数字。

*/

func minimumOperations(num string) int {

	//25,50,100,125

	pattern := []string{"25", "50", "75", "00"}

	count := strings.Count(num, "0")
	n := len(num)
	ans := n
	if count > 0 {
		ans = n - 1
	}
	for _, v := range pattern {
		cur := 0
		p := len(v) - 1
		for i := n - 1; i >= 0 && p >= 0; i-- {
			if v[p] == num[i] {
				p--
			} else {
				cur++
			}
		}
		if p < 0 {
			ans = min(ans, cur)
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumOperations("2245047"))
}
