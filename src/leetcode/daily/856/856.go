package main

import "fmt"

/*
*
给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

() 得 1 分。
AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
(A) 得 2 * A 分，其中 A 是平衡括号字符串。

//(()(()))
*/
func scoreOfParentheses(s string) int {

	ans := 0

	cnt := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
			if s[i-1] == '(' {
				ans += 1 << cnt
			}
		}
	}

	return ans
}

func scoreOfParentheses2(s string) int {

	st := []int{0}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, 0)
		} else {
			cur := st[len(st)-1]
			st = st[:len(st)-1]
			st[len(st)-1] += max(cur*2, 1)
		}
	}
	return st[0]
}

func main() {
	fmt.Println(scoreOfParentheses2("(())"))
}
