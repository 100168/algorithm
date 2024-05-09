package main

import "fmt"

/*
*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/
func minWindow(s string, t string) string {

	n := len(s)
	cntT := make(map[uint8]int)
	cntS := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		cntT[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		cntS[s[i]]++
	}
	for k, v := range cntT {
		if cntS[k] < v {
			return ""
		}
	}

	l := 0

	ans := n

	ll, rr := 0, n-1
	w := make(map[uint8]int)
next:
	for r := 0; r < len(s); r++ {
		w[s[r]]++
		//窗口内的数量大于需要数量，左边界往右滑
		for w[s[l]] > cntT[s[l]] {
			w[s[l]]--
			l++
		}
		for k, v := range cntT {
			if w[k] < v {
				continue next
			}
		}
		if ans > r-l+1 {
			ll = l
			rr = r
		}
		ans = min(ans, r-l+1)

	}
	return s[ll : rr+1]
}

func main() {
	fmt.Println(minWindow("aa", "aa"))
}
