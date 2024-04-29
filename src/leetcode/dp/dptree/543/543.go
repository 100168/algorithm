package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。
*/
func diameterOfBinaryTree(root *TreeNode) int {

	ans := 0
	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)
		ans = max(ans, leftMax+rightMax)
		return max(leftMax, rightMax) + 1
	}
	dfs(root)
	return ans
}
