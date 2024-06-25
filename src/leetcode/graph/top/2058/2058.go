package main

import "fmt"

/*
*
给你一个整数 n ，表示有 n 节课，课程编号从 1 到 n 。
同时给你一个二维整数数组 relations ，
其中 relations[j] = [prevCoursej, nextCoursej] ，
表示课程 prevCoursej 必须在课程 nextCoursej 之前 完成（先修课的关系）。
同时给你一个下标从 0 开始的整数数组 time ，其中 time[i] 表示完成第 (i+1) 门课程需要花费的 月份 数。

请你根据以下规则算出完成所有课程所需要的 最少 月份数：

如果一门课的所有先修课都已经完成，你可以在 任意 时间开始这门课程。
你可以 同时 上 任意门课程 。
请你返回完成所有课程所需要的 最少 月份数。

注意：测试数据保证一定可以完成所有课程（也就是先修课的关系构成一个有向无环图）。

1.无环有向图，直接dfs dp 简称树形dp
*/
func minimumTime(n int, relations [][]int, time []int) int {
	g := make([][]int, n+1)
	for _, v := range relations {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
	}
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int

	dfs = func(x int) int {

		if memo[x] != -1 {
			return memo[x]
		}

		cur := 0

		for _, y := range g[x] {
			cur = max(cur, dfs(y))
		}
		cur += time[x-1]
		memo[x] = cur
		return cur
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dfs(i))
	}
	return ans

}

/*
*

top排序+dp还是挺牛逼的
*/
func minimumTime2(n int, relations [][]int, time []int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, r := range relations {
		x, y := r[0]-1, r[1]-1
		g[x] = append(g[x], y) // 建图
		deg[y]++               // 可以理解为 y 的先修课的个数
	}

	var q []int
	for i, d := range deg {
		if d == 0 { // 没有先修课
			q = append(q, i)
		}
	}
	f := make([]int, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		// x 出队，意味着 x 的所有先修课都上完了
		f[x] += time[x] // 加上当前课程的时间，就得到了最终的 f[x]
		ans = max(ans, f[x])
		for _, y := range g[x] { // 遍历 x 的邻居 y
			f[y] = max(f[y], f[x])     // 更新 f[y] 的所有先修课程耗时的最大值
			if deg[y]--; deg[y] == 0 { // y 的先修课已上完
				q = append(q, y)
			}
		}
	}
	return
}
func main() {

	fmt.Println(minimumTime(3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}))
}
