package main

import (
	"fmt"
	"math"
)

/**
给你一个长度为 n 的字符串 source ，一个字符串 pattern 且它是 source 的
子序列
 ，和一个 有序 整数数组 targetIndices ，整数数组中的元素是 [0, n - 1] 中 互不相同 的数字。

定义一次 操作 为删除 source 中下标在 idx 的一个字符，且需要满足：

idx 是 targetIndices 中的一个元素。
删除字符后，pattern 仍然是 source 的一个
子序列
 。
执行操作后 不会 改变字符在 source 中的下标位置。比方说，如果从 "acb" 中删除 'c' ，下标为 2 的字符仍然是 'b' 。

请你返回 最多 可以进行多少次删除操作。

子序列指的是在原字符串里删除若干个（也可以不删除）字符后，不改变顺序地连接剩余字符得到的字符串。



示例 1：

输入：source = "abbaa", pattern = "aba", targetIndices = [0,1,2]

输出：1

解释：

不能删除 source[0] ，但我们可以执行以下两个操作之一：

删除 source[1] ，source 变为 "a_baa" 。
删除 source[2] ，source 变为 "ab_aa" 。
示例 2：

输入：source = "bcda", pattern = "d", targetIndices = [0,3]

输出：2

解释：

进行两次操作，删除 source[0] 和 source[3] 。

示例 3：

输入：source = "dda", pattern = "dda", targetIndices = [0,1,2]

输出：0

解释：

不能在 source 中删除任何字符。

示例 4：

输入：source = "yeyeykyded", pattern = "yeyyd", targetIndices = [0,2,3,4]

输出：2

解释：

进行两次操作，删除 source[2] 和 source[3] 。



提示：

1 <= n == source.length <= 3 * 103
1 <= pattern.length <= n
1 <= targetIndices.length <= n
targetIndices 是一个升序数组。
输入保证 targetIndices 包含的元素在 [0, n - 1] 中且互不相同。
source 和 pattern 只包含小写英文字母。
输入保证 pattern 是 source 的一个子序列。
*/

func maxRemovals(source string, pattern string, targetIndices []int) int {
	n := len(pattern)
	canDel := make(map[int]bool)
	for _, v := range targetIndices {
		canDel[v] = true
	}

	m := len(source)
	f := make([][]int, n)

	for i := range f {

		f[i] = make([]int, m)

		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(p1 int, p2 int) int {

		cur := math.MinInt / 2
		if p1 < 0 {
			cur = 0
			for ; p2 >= 0; p2-- {
				if canDel[p2] {
					cur++
				}
			}
			return cur
		}
		if p2 < 0 {
			return math.MinInt / 2
		}

		if f[p1][p2] != -1 {
			return f[p1][p2]
		}

		if source[p2] == pattern[p1] {
			cur = dfs(p1-1, p2-1)
		}
		if canDel[p2] {
			cur = max(cur, dfs(p1, p2-1)+1)
		} else {
			cur = max(cur, dfs(p1, p2-1))
		}

		f[p1][p2] = cur
		return cur
	}
	return dfs(n-1, m-1)
}

func main() {
	fmt.Println(maxRemovals("abbaa", "aba", []int{0, 1, 2}))
}
