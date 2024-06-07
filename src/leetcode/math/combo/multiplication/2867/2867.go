package main

import "fmt"

/*
*
给你一棵 n 个节点的无向树，节点编号为 1 到 n 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，
其中 edges[i] = [ui, vi] 表示节点 ui 和 vi 在树中有一条边。
请你返回树中的 合法路径数目 。
如果在节点 a 到节点 b 之间 恰好有一个 节点的编号是质数，那么我们称路径 (a, b) 是 合法的 。
注意：
路径 (a, b) 指的是一条从节点 a 开始到节点 b 结束的一个节点序列，序列中的节点 互不相同 ，且相邻节点之间在树上有一条边。
路径 (a, b) 和路径 (b, a) 视为 同一条 路径，且只计入答案 一次 。

思路：
1.预处理所有质数（质数筛）
2.树形dp需要根据当前数是否是质数进行分类讨论
*/
func countPaths(n int, edges [][]int) int64 {

	primes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i <= n; i++ {

		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}
	ans := 0

	g := make([][]int, n+1)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	//左边包含一个质数，左边包含0个质数

	var dfs func(int, int) (int, int)

	dfs = func(x, fa int) (int, int) {
		cntOne := 0
		cntZero := 0
		//当前数为质数需要
		if primes[x] {
			cntOne = 1
		} else {
			cntZero = 1
		}
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			curOne, curZero := dfs(y, x)
			//当前数是质数
			if primes[x] {
				ans += cntOne * curZero
			} else {
				ans += cntZero*curOne + cntOne*curZero
			}
			cntZero += curZero
			if !primes[x] {
				cntOne += curOne
			} else {
				cntOne += curZero
			}

		}
		if primes[x] {
			//左0+右0 +自己1= cntZero+1
			return cntOne, 0
		}
		return cntOne, cntZero
	}

	dfs(1, -1)
	return int64(ans)
}

func main() {
	fmt.Println(countPaths(5, [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}}))
}
