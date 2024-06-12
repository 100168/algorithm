package main

import (
	"fmt"
	"strings"
)

/*
*
给你一个字符串 s ，它包含一个或者多个单词。单词之间用单个空格 ' ' 隔开。

如果字符串 t 中第 i 个单词是 s 中第 i 个单词的一个 排列 ，
那么我们称字符串 t 是字符串 s 的同位异构字符串。

比方说，"acb dfe" 是 "abc def" 的同位异构字符串，但是 "def cab" 和 "adc bef" 不是。
请你返回 s 的同位异构字符串的数目，由于答案可能很大，请你将它对 109 + 7 取余 后返回。
*/

const mod = int(1e9 + 7)

func countAnagrams(s string) int {

	split := strings.Split(s, " ")

	maxLen := 0
	for _, v := range split {
		maxLen = max(maxLen, len(v))
	}

	comb := make([]int, maxLen+1)
	comb[1] = 1
	ans := 1

	div := 1

	for i := 2; i <= maxLen; i++ {

		comb[i] = comb[i-1] * i % mod
	}
	for _, v := range split {

		cntMap := make(map[rune]int)
		ans = ans * comb[len(v)] % mod
		for i := range v {
			cntMap[rune(v[i])]++
		}

		for _, v := range cntMap {
			div = div * comb[v] % mod
		}

	}
	return ans * pow(div, mod-2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func main() {
	fmt.Println(countAnagrams("acb dfe"))
	fmt.Println(countAnagrams("too hot"))
}
