package main

/*
如果字符串满足以下条件之一，则可以称之为 有效括号字符串（valid parentheses string，可以简写为 VPS）：

字符串是一个空字符串 ""，或者是一个不为 "(" 或 ")" 的单字符。
字符串可以写为 AB（A 与 B字符串连接），其中 A 和 B 都是 有效括号字符串 。
字符串可以写为 (A)，其中 A 是一个 有效括号字符串 。
类似地，可以定义任何有效括号字符串p 的 嵌套深度 depth(S)：

depth("") = 0
depth(C) = 0，其中 C 是单个字符的字符串，且该字符不是 "(" 或者 ")"
depth(A + B) = max(depth(A), depth(B))，其中 A 和 B 都是 有效括号字符串
depth("(" + A + ")") = 1 + depth(A)，其中 A 是一个 有效括号字符串
例如：""、"()()"、"()(()())" 都是 有效括号字符串（嵌套深度分别为 0、1、2），而 ")(" 、"(()" 都不是 有效括号字符串 。

给你一个 有效括号字符串 s，返回该字符串的 s 嵌套深度 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-nesting-depth-of-the-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxDepth(s string) int {

	n := len(s)
	des := []byte(s)
	cnt := 0
	ans := 0
	for i := 0; i < n; i++ {

		if des[i] == byte('(') {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		}
		if des[i] == byte(')') {
			cnt--
			if cnt < 0 {
				return 0
			}
		}
	}

	if cnt == 0 {
		return ans
	}
	return 0
}
