package main

import "math"

/*
*

给你一个字符串 s ，你需要将它分割成一个或者更多的 平衡 子字符串。
比方说，s == "ababcc" 那么 ("abab", "c", "c") ，("ab", "abc", "c") 和 ("ababcc") 都是合法分割，
但是 ("a", "bab", "cc") ，("aba", "bc", "c") 和 ("ab", "abcc") 不是，不平衡的子字符串用粗体表示。

请你返回 s 最少 能分割成多少个平衡子字符串。

注意：一个 平衡 字符串指的是字符串中所有字符出现的次数都相同。
*/
func minimumSubstringsInPartition(s string) int {

	n := len(s)

	memo := make([]int, n)

	for i := range memo {

		memo[i] = -1
	}
	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		w := make(map[uint8]int)

		if memo[i] != -1 {
			return memo[i]
		}
		cur := math.MaxInt
	next:
		for j := i; j >= 0; j-- {
			w[s[j]]++
			pre := -1
			for _, v := range w {
				if pre != -1 && v != pre {
					continue next
				}
				pre = v
			}
			cur = min(cur, dfs(j-1)+1)
		}
		memo[i] = cur
		return cur
	}
	return dfs(n - 1)
}

func main() {
	println(minimumSubstringsInPartition("a,b,c,c"))
}
