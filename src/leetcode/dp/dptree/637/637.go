package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。

两个节点之间的路径长度 由它们之间的边数表示。*/

func longestUnivaluePath(root *TreeNode) int {

	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(root *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		cur := 1

		maxVal := 1
		l, r := dfs(node.Left), dfs(node.Right)

		if l > 0 && node.Val == node.Left.Val {
			maxVal = max(maxVal, l+1)
			cur += l
		}
		if r > 0 && node.Val == node.Right.Val {
			cur += r
			maxVal = max(maxVal, r+1)
		}
		ans = max(ans, cur-1)
		return maxVal
	}
	dfs(root)
	return ans
}
