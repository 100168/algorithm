package main

import (
	"fmt"
	"math"
)

/*
给你一个字符串 s ，它每一位都是 1 到 9 之间的数字组成，同时给你一个整数 k 。

如果一个字符串 s 的分割满足以下条件，我们称它是一个 好 分割：

s 中每个数位 恰好 属于一个子字符串。
每个子字符串的值都小于等于 k 。
请你返回 s 所有的 好 分割中，子字符串的 最少 数目。如果不存在 s 的 好 分割，返回 -1 。

注意：

一个字符串的 值 是这个字符串对应的整数。比方说，"123" 的值为 123 ，"1" 的值是 1 。
子字符串 是字符串中一段连续的字符序列。
*/
var dp []int

// dfs
// 1.从左往右
func minimumPartition(s string, k int) int {
	if k < 10 {
		for i := range s {
			if int(s[i]-'0') > k {
				return -1
			}
		}
	}

	dp = make([]int, len(s)+1)
	dp[0] = 0

	for i := range dp {
		dp[i] = -1
	}
	return dfs(0, s, k)

}
func dfs(index int, s string, k int) int {

	if index >= len(s) {
		dp[index] = 0
		return 0
	}

	cur := 0
	cnt := math.MaxInt / 2
	curIndex := index
	for index < len(s) {
		cur = cur*10 + int(s[index]-'0')
		if cur > k {
			break
		}
		index++
		if dp[index] != -1 {
			cnt = min(dp[index]+1, cnt)
		} else {
			cnt = min(dfs(index, s, k)+1, cnt)
		}
	}
	dp[curIndex] = cnt
	return cnt
}

func minimumPartition2(s string, k int) int {
	pre := 0
	ans := 0
	for _, v := range s {
		c := int(v - '0')
		if pre*10+c <= k {
			pre = pre*10 + c
			continue
		}
		if c > k {
			return -1
		}
		pre = c
		ans++
	}
	if pre > 0 {
		ans++
	}
	return ans

}

func main() {
	fmt.Println(minimumPartition("165462", 60))
	fmt.Println(minimumPartition2("165462", 60))
}
