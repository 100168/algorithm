package main

import (
	"fmt"
)

/*
*
给你一棵 n 个节点且根节点为编号 0 的树，节点编号为 0 到 n - 1 。
这棵树用一个长度为 n 的数组 parent 表示，其中 parent[i] 是第 i 个节点的父亲节点的编号。
由于节点 0 是根，parent[0] == -1 。

给你一个长度为 n 的字符串 s ，其中 s[i] 是节点 i 对应的字符。

对于节点编号从 1 到 n - 1 的每个节点 x ，我们 同时 执行以下操作 一次 ：

找到距离节点 x 最近 的祖先节点 y ，且 s[x] == s[y] 。
如果节点 y 不存在，那么不做任何修改。
否则，将节点 x 与它父亲节点之间的边 删除 ，在 x 与 y 之间连接一条边，使 y 变为 x 新的父节点。
请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是 最终 树中，节点 i 为根的子树的 大小 。
*/
func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)

	g := make([][]int, n)
	for i, v := range parent {
		if i == 0 {
			continue
		}
		g[v] = append(g[v], i)
	}
	ans := make([]int, n)

	charMap := make([]int, 26)

	for i := range charMap {
		charMap[i] = -1
	}

	var reboot func(int)

	reboot = func(x int) {
		pre := charMap[s[x]-'a']
		if pre >= 0 {
			parent[x] = pre
		}
		charMap[s[x]-'a'] = x
		for _, y := range g[x] {
			reboot(y)
		}
		charMap[s[x]-'a'] = pre
	}

	reboot(0)

	g = make([][]int, n)
	for i, v := range parent {
		if i == 0 {
			continue
		}
		g[v] = append(g[v], i)
	}
	var dfs func(int) int

	dfs = func(x int) int {
		c := 1
		for _, v := range g[x] {
			c += dfs(v)
		}
		ans[x] = c
		return c
	}

	dfs(0)
	return ans

}
func main() {
	fmt.Println(findSubtreeSizes([]int{1, 0, 0, 1, 1, 1}, "abaabc"))
	fmt.Println(findSubtreeSizes([]int{-1, 10, 0, 12, 10, 18, 11, 12, 2, 3, 2, 2, 2, 0, 4, 11, 4, 2, 0}, "babadabbdabcbaceeda"))
}
