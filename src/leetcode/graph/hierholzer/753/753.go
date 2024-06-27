package main

import (
	"math"
	"strconv"
)

/*
*
有一个需要密码才能打开的保险箱。密码是 n 位数, 密码的每一位都是范围 [0, k - 1] 中的一个数字。

保险箱有一种特殊的密码校验方法，你可以随意输入密码序列，保险箱会自动记住 最后 n 位输入 ，如果匹配，则能够打开保险箱。

例如，正确的密码是 "345" ，并且你输入的是 "012345" ：
输入 0 之后，最后 3 位输入是 "0" ，不正确。
输入 1 之后，最后 3 位输入是 "01" ，不正确。
输入 2 之后，最后 3 位输入是 "012" ，不正确。
输入 3 之后，最后 3 位输入是 "123" ，不正确。
输入 4 之后，最后 3 位输入是 "234" ，不正确。
输入 5 之后，最后 3 位输入是 "345" ，正确，打开保险箱。
在只知道密码位数 n 和范围边界 k 的前提下，请你找出并返回确保在输入的 某个时刻 能够打开保险箱的任一 最短 密码序列 。

f[0][1] f[0][2] f[0][3]
*/
func crackSafe(n int, k int) string {
	seen := map[int]bool{}
	ans := ""
	highest := int(math.Pow(10, float64(n-1)))

	var dfs func(int)
	dfs = func(node int) {
		for x := 0; x < k; x++ {
			nei := node*10 + x
			if !seen[nei] {
				seen[nei] = true
				dfs(nei % highest)
				ans += strconv.Itoa(x)
			}
		}
	}
	dfs(0)
	for i := 1; i < n; i++ {
		ans += "0"
	}
	return ans
}
