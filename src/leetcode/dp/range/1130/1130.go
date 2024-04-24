package main

import (
	"math"
	"slices"
)

/*
*
给你一个正整数数组 arr，考虑所有满足以下条件的二叉树：

每个节点都有 0 个或是 2 个子节点。
数组 arr 中的值与树的中序遍历中每个叶节点的值一一对应。
每个非叶节点的值等于其左子树和右子树中叶节点的最大值的乘积。
在所有这样的二叉树中，返回每个非叶节点的值的最小可能总和。这个和的值是一个 32 位整数。

如果一个节点有 0 个子节点，那么该节点为叶节点。
*/
func mctFromLeafValues(arr []int) int {

	n := len(arr)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(l int, r int) int {

		if l >= r {
			return 0
		}
		if memo[l][r] != -1 {
			return memo[l][r]
		}
		cur := math.MaxInt / 2
		for i := l; i < r; i++ {
			lMax := slices.Max(arr[l : i+1])
			rMax := slices.Max(arr[i+1 : r+1])
			curVal := lMax * rMax
			cur = min(cur, dfs(l, i)+dfs(i+1, r)+curVal)
		}
		memo[l][r] = cur
		return cur
	}

	return dfs(0, n-1)
}

func main() {
	println(mctFromLeafValues([]int{6, 2, 4}))
}
