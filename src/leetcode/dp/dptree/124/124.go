package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/
func maxPathSum(root *TreeNode) int {

	ans := math.MinInt
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		//可以左右同时选
		curMax := max(node.Val, node.Val+left, node.Val+right, node.Val+left+right)
		ans = max(ans, curMax)
		//不能左右同时选
		return max(node.Val, node.Val+left, node.Val+right)
	}
	dfs(root)
	return ans
}
