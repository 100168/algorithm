package main

import "fmt"

/**

一个括号字符串是只由 '(' 和 ')' 组成的 非空 字符串。如果一个字符串满足下面 任意 一个条件，那么它就是有效的：

字符串为 ().
它可以表示为 AB（A 与 B 连接），其中A 和 B 都是有效括号字符串。
它可以表示为 (A) ，其中 A 是一个有效括号字符串。
给你一个括号字符串 s 和一个字符串 locked ，两者长度都为 n 。locked 是一个二进制字符串，只包含 '0' 和 '1' 。对于 locked 中 每一个 下标 i ：

如果 locked[i] 是 '1' ，你 不能 改变 s[i] 。
如果 locked[i] 是 '0' ，你 可以 将 s[i] 变为 '(' 或者 ')' 。
如果你可以将 s 变为有效括号字符串，请你返回 true ，否则返回 false 。


输入：s = "))()))", locked = "010100"
输出：true
解释：locked[1] == '1' 和 locked[3] == '1' ，所以我们无法改变 s[1] 或者 s[3] 。
我们可以将 s[0] 和 s[4] 变为 '(' ，不改变 s[2] 和 s[5] ，使 s 变为有效字符串。

*/

func canBeValid(s string, locked string) bool {

	n := len(s)
	if n%2 != 0 {
		return false
	}
	var stLeft []int

	var stChange []int

	for i := range locked {
		if locked[i] == '1' {
			if s[i] == '(' {
				stLeft = append(stLeft, i)
			} else {
				if len(stLeft) == 0 && len(stChange) == 0 {
					return false
				}
				if len(stLeft) > 0 {
					stLeft = stLeft[:len(stLeft)-1]
					continue
				}
				stChange = stChange[:len(stChange)-1]
			}
		} else {
			stChange = append(stChange, i)
		}
	}

	for i := len(stLeft) - 1; i >= 0; i-- {

		l := stLeft[i]

		if len(stChange) == 0 {
			return false
		}
		change := stChange[len(stChange)-1]
		stChange = stChange[:len(stChange)-1]

		if change < l {
			return false
		}
	}
	return true
}

func canBeValid2(s string, locked string) bool {

	cnt := 0
	n := len(s)
	if n%2 != 0 {
		return false
	}
	cntOps := 0
	for i := range locked {
		if locked[i] == '1' {
			if s[i] == '(' {
				cnt++
			} else {
				cnt--
				if cnt < 0 {
					cntOps--
					if cntOps < 0 {
						return false
					}
					cnt = 0
				}
			}
		} else {
			cntOps++
		}
	}

	cntOps = 0
	cnt = 0

	for i := n - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				cnt++
			} else {
				cnt--
				if cnt < 0 {
					cntOps--
					if cntOps < 0 {
						return false
					}
					cnt = 0
				}
			}
		} else {
			cntOps++
		}
	}
	return true
}

func main() {
	fmt.Println(canBeValid2("((()(()()))()((()()))))()((()(()", "10111100100101001110100010001001"))
	fmt.Println(canBeValid2("())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(", "100011110110011011010111100111011101111110000101001101001111"))
	fmt.Println(canBeValid2("))))(())((()))))((()((((((())())((()))((((())()()))(()", "101100101111110000000101000101001010110001110000000101"))
	fmt.Println(canBeValid2("())()))()(()(((())(()()))))((((()())(())", "1011101100010001001011000000110010100101"))
}
