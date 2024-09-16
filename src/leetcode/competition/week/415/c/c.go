package main

import (
	"math"
)

/*
*
给你一个字符串数组 words 和一个字符串 target。

如果字符串 x 是 words 中 任意 字符串的
前缀
，则认为 x 是一个 有效 字符串。

现计划通过 连接 有效字符串形成 target ，请你计算并返回需要连接的 最少 字符串数量。如果无法通过这种方式形成 target，则返回 -1。

示例 1：

输入： words = ["abc","aaaaa","bcdef"], target = "aabcdabc"

输出： 3

解释：

target 字符串可以通过连接以下有效字符串形成：

words[1] 的长度为 2 的前缀，即 "aa"。
words[2] 的长度为 3 的前缀，即 "bcd"。
words[0] 的长度为 3 的前缀，即 "abc"。


*/

type trie struct {
	root *node
}

type node struct {
	next [26]*node
	path int
	end  int
}

func (t *trie) insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.next[c] == nil {
			cur.next[c] = &node{}
		}
		cur = cur.next[c]
		cur.path++
	}
	cur.end++
}
func (t *trie) find(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		c := int(word[i] - 'a')
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}
	return cur.path > 0
}
func minValidStrings(words []string, target string) int {

	n := len(target)

	t := new(trie)
	t.root = new(node)

	for _, v := range words {
		t.insert(v)
	}

	f := make([]int, n)

	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int

	dfs = func(i int) int {

		if i >= n {

			return 0
		}

		if f[i] != -1 {
			return f[i]
		}
		cur := math.MaxInt / 2

		for j := i; j < n; j++ {
			if t.find(target[i : j+1]) {
				cur = min(cur, dfs(j-1)+1)
			} else {
				break
			}
		}

		f[i] = cur
		return cur

	}
	ans := dfs(n - 1)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func minValidStrings2(words []string, target string) int {

	n := len(target)

	maxL := 0
	for _, v := range words {

		maxL = max(maxL, len(v))
	}

	preMap := make([]map[int]bool, maxL)

	for i := range preMap {
		preMap[i] = make(map[int]bool)
	}
	base := 131313
	mod := int(1e9 + 7)

	for _, v := range words {

		p := 0

		for i := range v {
			p = (p*base + int(v[i]-'a')) % mod
			preMap[i][p] = true
		}
	}

	f := make([]int, n)

	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int

	dfs = func(i int) int {

		if i >= n {

			return 0
		}

		if f[i] != -1 {
			return f[i]
		}
		cur := math.MaxInt / 2

		p := 0
		for j := i; j < n && j-i+1 <= maxL; j++ {
			p = (p*base + int(target[j]-'a')) % mod
			if preMap[j-i][p] {
				cur = min(cur, dfs(j+1)+1)
			}
		}

		f[i] = cur
		return cur

	}
	ans := dfs(0)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func main() {

}
