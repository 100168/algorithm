package main

import (
	"fmt"
	"math"
	"strings"
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

	orr := make([][]int, n)
	for i := range orr {
		orr[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		s1 := words[i]
		for j := i + 1; j < n; j++ {
			s2 := words[j]
			for k := 1; k <= min(len(s1), len(s2)); k++ {
				if s1[len(s1)-k:] == s2[:k] {
					orr[i][j] = k
					orr[j][i] = k
				}
			}
		}
	}

	memo := make([][]int, 1<<n)
	parent := make([][]int, 1<<n)

	for i := range memo {
		memo[i] = make([]int, n)
		parent[i] = make([]int, n)

		for j := range memo[i] {
			memo[i][j] = -1
			parent[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(mask int, pre int) int {

		if mask == 1<<n-1 {
			return 0
		}

		if memo[mask][pre] != -1 {
			return memo[mask][pre]
		}

		cur := math.MinInt
		for i := 0; i < n; i++ {
			if 1<<i&mask == 0 {
				cur = min(cur, dfs(mask|1<<i, i)+orr[pre][mask])
			}
		}
	}

	dfs(0, -1)
	return ans
}

func merge(s1, s2 string) int {

	n1 := len(s1)
	n2 := len(s2)

	index := 1

	maxI := 0
	for index <= min(n1, n2) {
		if s1[n1-index:] == s2[:index] {
			maxI = index
		}
		index++
	}
	return maxI

}

func check(s string, words []string) bool {
	for _, v := range words {
		if strings.Index(s, v) < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(shortestSuperstring([]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}))
	fmt.Println(merge("catg", "atg"))
}
