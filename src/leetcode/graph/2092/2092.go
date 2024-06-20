package main

import (
	"sort"
)

/*
*给你一个整数 n ，表示有 n 个专家从 0 到 n - 1 编号。另外给你一个下标从 0 开始的二维整数数组 meetings ，
其中 meetings[i] = [xi, yi, timei] 表示专家 xi 和专家 yi 在时间 timei 要开一场会。一个专家可以同时参加 多场会议 。
最后，给你一个整数 firstPerson 。

专家 0 有一个 秘密 ，最初，他在时间 0 将这个秘密分享给了专家 firstPerson 。
接着，这个秘密会在每次有知晓这个秘密的专家参加会议时进行传播。
更正式的表达是，每次会议，如果专家 xi 在时间 timei 时知晓这个秘密，那么他将会与专家 yi 分享这个秘密，反之亦然。

秘密共享是 瞬时发生 的。也就是说，在同一时间，一个专家不光可以接收到秘密，还能在其他会议上与其他专家分享。

在所有会议都结束之后，返回所有知晓这个秘密的专家列表。你可以按 任何顺序 返回答案。

输入：n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1
输出：[0,1,2,3,5]
解释：
时间 0 ，专家 0 将秘密与专家 1 共享。
时间 5 ，专家 1 将秘密与专家 2 共享。
时间 8 ，专家 2 将秘密与专家 3 共享。
时间 10 ，专家 1 将秘密与专家 5 共享。
因此，在所有会议结束后，专家 0、1、2、3 和 5 都将知晓这个秘密。

思路：
我是废物
*/
func findAllPeople(_ int, meetings [][]int, firstPerson int) (ans []int) {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] }) // 按照时间排序

	haveSecret := map[int]bool{0: true, firstPerson: true} // 一开始 0 和 firstPerson 都知道秘密
	for i, m := 0, len(meetings); i < m; {
		g := map[int][]int{}
		time := meetings[i][2]
		// 遍历时间相同的会议。注意这里的 i 和外层循环的 i 是同一个变量，所以整个循环部分的时间复杂度是线性的
		for ; i < m && meetings[i][2] == time; i++ {
			v, w := meetings[i][0], meetings[i][1]
			g[v] = append(g[v], w) // 建图
			g[w] = append(g[w], v)
		}

		vis := map[int]bool{} // 避免重复访问专家
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			haveSecret[v] = true
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		for v := range g {
			if haveSecret[v] && !vis[v] { // 从在图上且知道秘密的专家出发，DFS 标记所有能到达的专家
				dfs(v)
			}
		}
	}
	for i := range haveSecret {
		ans = append(ans, i) // 注意可以按任何顺序返回答案
	}
	return
}
