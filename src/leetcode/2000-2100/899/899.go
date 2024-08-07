package main

import (
	"fmt"
	"slices"
)

/*
*
给定一个字符串 s 和一个整数 k 。你可以从 s 的前 k 个字母中选择一个，并把它加到字符串的末尾。
返回 在应用上述步骤的任意数量的移动后，字典序最小的字符串 。

输入：s = "cbda", k = 2    bddaac
输出："acb"
解释：
在第一步中，我们将第一个字符（“c”）移动到最后，获得字符串 “bac”。
在第二步中，我们将第一个字符（“b”）移动到最后，获得最终结果 “acb”。

输入：s = "aaacb", k = 3
输出："acbaa" aabac   aaabc
解释：
在第一步中，我们将第一个字符（“b”）移动到最后，获得字符串 “aacab”。
在第二步中，我们将第三个字符（“c”）移动到最后，获得最终结果 “aaabc”。
*/
func orderlyQueue(s string, k int) string {

	bytes := []byte(s)

	n := len(bytes)
	ans := s
	if k == 1 {
		for i := 1; i < n; i++ {
			cur := s[i:] + s[:i]
			if ans > cur {
				ans = cur
			}
		}
		return ans
	}

	slices.Sort(bytes)
	return string(bytes)

}

func main() {
	fmt.Println(orderlyQueue("cba", 1))
}
