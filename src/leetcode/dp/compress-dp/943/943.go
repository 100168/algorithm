package main

import (
	"fmt"
	"math"
)

/**
给定一个字符串数组 words，找到以 words 中每个字符串作为子字符串的最短字符串。如果有多个有效最短字符串满足题目条件，返回其中 任意一个 即可。

我们可以假设 words 中没有字符串是 words 中另一个字符串的子字符串。*/

func shortestSuperstring(words []string) string {

	n := len(words)

	ans := ""
	for _, word := range words {
		ans += word
	}

	overLap := make([][]int, n)
	for i := range overLap {
		overLap[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		s1 := words[i]
		for j := 0; j < n; j++ {
			s2 := words[j]
			for k := 1; k <= min(len(s1), len(s2)); k++ {
				if s1[len(s1)-k:] == s2[:k] {
					overLap[i][j] = k
				}
			}
		}
	}
	memo := make([]map[int]string, 1<<n)

	for i := range memo {
		memo[i] = make(map[int]string, n)
	}
	var dfs func(int, int) string

	dfs = func(mask int, pre int) string {

		if mask == 1<<n-1 {
			return words[pre]
		}

		if _, ok := memo[mask][pre]; ok {
			return memo[mask][pre]
		}

		curLen := math.MaxInt / 2
		cur := ""

		for i := 0; i < n; i++ {
			if 1<<i&mask == 0 {
				overlap := overLap[pre][i]
				preL := len(words[pre])
				next := words[pre][:preL-overlap] + dfs(mask|1<<i, i)
				if len(next) < curLen {
					cur = next
					curLen = len(next)
				}
			}
		}

		memo[mask][pre] = cur
		return cur
	}
	for i := 0; i < n; i++ {
		cur := dfs(1<<i, i)
		if len(cur) < len(ans) {
			ans = cur
		}
	}
	return ans
}

func main() {
	fmt.Println(shortestSuperstring([]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}))
	//fmt.Println(merge("catg", "atg"))
}
