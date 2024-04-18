package main

import "math"

/*
给你一棵 n 个节点的 无向 树，节点从 0 到 n - 1 编号。树以长度为 n - 1 下标从 0 开始的二维整数数组 edges 的形式给你，其中 edges[i] = [ui, vi] 表示树中节点 ui 和 vi 之间有一条边。同时给你一个 正 整数 k 和一个长度为 n 下标从 0 开始的 非负 整数数组 nums ，其中 nums[i] 表示节点 i 的 价值 。

Alice 想 最大化 树中所有节点价值之和。为了实现这一目标，Alice 可以执行以下操作 任意 次（包括 0 次）：

选择连接节点 u 和 v 的边 [u, v] ，并将它们的值更新为：
nums[u] = nums[u] XOR k
nums[v] = nums[v] XOR k
请你返回 Alice 通过执行以上操作 任意次 后，可以得到所有节点 价值之和 的 最大值 。
*/
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		//f0 当前为偶数，f1当前为奇数
		f0, f1 := 0, math.MinInt
		for _, y := range g[x] {
			if y != fa {
				r0, r1 := dfs(y, x)
				f0, f1 = max(f0+r0, f1+r1), max(f1+r0, f0+r1)
			}
		}
		return max(f0+nums[x], f1+(nums[x]^k)), max(f1+nums[x], f0+(nums[x]^k))
	}
	ans, _ := dfs(0, -1)
	return int64(ans)
}

func main() {
	println(maximumValueSum([]int{1, 2, 1}, 3, [][]int{{0, 1}, {0, 2}}))
}

// 下班搞
func maximumValueSum2(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		f0, f1 := 0, math.MinInt // f[x][0] 和 f[x][1]
		for _, y := range g[x] {
			if y != fa {
				r0, r1 := dfs(y, x)
				f0, f1 = max(f0+r0, f1+r1), max(f1+r0, f0+r1)
			}
		}
		return max(f0+nums[x], f1+(nums[x]^k)), max(f1+nums[x], f0+(nums[x]^k))
	}
	ans, _ := dfs(0, -1)
	return int64(ans)
}

func maximumValueSum3(nums []int, k int, _ [][]int) int64 {
	f0, f1 := 0, math.MinInt
	for _, x := range nums {
		f0, f1 = max(f0+x, f1+(x^k)), max(f1+x, f0+(x^k))
	}
	return int64(f0)
}
