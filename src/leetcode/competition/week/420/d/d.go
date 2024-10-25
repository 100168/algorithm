package main

import (
	"fmt"
	"time"
)

/*
*
给你一棵 n 个节点的树，树的根节点为 0 ，n 个节点的编号为 0 到 n - 1 。
这棵树用一个长度为 n 的数组 parent 表示，其中 parent[i] 是节点 i 的父节点。由于节点 0 是根节点，所以 parent[0] == -1 。

给你一个长度为 n 的字符串 s ，其中 s[i] 是节点 i 对应的字符。

一开始你有一个空字符串 dfsStr ，定义一个递归函数 dfs(int x) ，它的输入是节点 x ，并依次执行以下操作：

按照 节点编号升序 遍历 x 的所有孩子节点 y ，并调用 dfs(y) 。
将 字符 s[x] 添加到字符串 dfsStr 的末尾。
注意，所有递归函数 dfs 都共享全局变量 dfsStr 。

你需要求出一个长度为 n 的布尔数组 answer ，对于 0 到 n - 1 的每一个下标 i ，你需要执行以下操作：

清空字符串 dfsStr 并调用 dfs(i) 。
如果结果字符串 dfsStr 是一个
回文串

	，answer[i] 为 true ，否则 answer[i] 为 false 。

请你返回字符串 answer 。
*/
func findAnswer(parent []int, s string) []bool {

	n := len(s)

	g := make([][]int, n)

	left := make([]int, n)
	right := make([]int, n)

	for i := 1; i < n; i++ {

		p := parent[i]

		g[p] = append(g[p], i)
	}
	ans := make([]bool, n)

	chars := make([]byte, n)
	var dfs func(x int)

	t := 0
	fmt.Println(time.Now())
	dfs = func(x int) {
		left[x] = t
		for _, y := range g[x] {
			dfs(y)
		}
		chars[t] = s[x] // 后序遍历
		t++
		right[x] = t
	}
	dfs(0)
	fmt.Println(time.Now())

	// 0123456789
	// #a#b#a#a#b#a#
	// abaaba

	p := longestPalindrome(chars)
	fmt.Println(time.Now())
	for i := 0; i < n; i++ {
		cur := right[i] + left[i]
		ans[i] = p[cur] >= right[i]-left[i]
	}
	return ans

}

func longestPalindrome(s []byte) []int {

	n := 2*len(s) + 1

	str := make([]byte, n)

	str[len(str)-1] = '#'

	for i, v := range s {
		str[i*2] = '#'
		str[i*2+1] = v
	}
	//每个点的最长回文半径
	p := make([]int, n)
	//c回文中心,r是回问半径（）
	c, r := 0, 0
	for i := 0; i < n; i++ {
		//自己也是回文串
		l := 1
		if r > i {
			l = min(r-i, p[2*c-i])
		}

		for i+l < n && i-l >= 0 && str[i+l] == str[i-l] {
			l++
		}

		if i+l > r {
			r = i + l
			c = i
		}
		p[i] = l
	}
	return p
}

func main() {
	fmt.Println(findAnswer([]int{-1, 0, 0, 1, 1, 2}, "zzzzz"))
}
