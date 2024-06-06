package main

import (
	"fmt"
	"math/bits"
)

/*
*
我们定义了一个函数 countUniqueChars(s) 来统计字符串 s 中的唯一字符，并返回唯一字符的个数。

例如：s = "LEETCODE" ，则其中 "L", "T","C","O","D" 都是唯一字符，因为它们只出现一次，所以 countUniqueChars(s) = 5 。

本题将会给你一个字符串 s ，我们需要返回 countUniqueChars(t) 的总和，其中 t 是 s 的子字符串。输入用例保证返回值为 32 位整数。

注意，某些子字符串可能是重复的，但你统计时也必须算上这些重复的子字符串（也就是说，你必须统计 s 的所有子字符串中的唯一字符）。
*/
func uniqueLetterString(s string) int {

	type pair struct {
		one, two, cnt int
	}
	var pre []pair

	ans := 0

	for i := range s {
		cur := int(s[i] - 'A')
		temp := pre
		pre = nil
		for _, v := range temp {
			if v.one&(1<<cur) == 0 && v.two&(1<<cur) == 0 {
				v.one ^= 1 << cur
			} else if v.one&(1<<cur) != 0 {
				v.one ^= 1 << cur
				v.two ^= 1 << cur
			}
			ans += bits.OnesCount(uint(v.one)) * v.cnt
			if len(pre) > 0 && pre[len(pre)-1].one == v.one && pre[len(pre)-1].two == v.two {
				pre[len(pre)-1].cnt += v.cnt
			} else {
				pre = append(pre, pair{v.one, v.two, v.cnt})
			}
		}
		pre = append(pre, pair{1 << cur, 0, 1})
		ans++
	}
	return ans
}

func main() {
	fmt.Println(uniqueLetterString("LEETCODE"))
}
