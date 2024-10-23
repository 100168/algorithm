package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
给你一棵 二叉树 的根节点 root 和一个整数k。

返回第 k 大的 完美二叉
子树
的大小，如果不存在则返回 -1。

完美二叉树 是指所有叶子节点都在同一层级的树，且每个父节点恰有两个子节点。

示例 1：

输入： root = [5,3,6,5,2,5,7,1,8,null,null,6,8], k = 2

输出： 3

解释：

完美二叉子树的根节点在图中以黑色突出显示。它们的大小按非递增顺序排列为 [3, 3, 1, 1, 1, 1, 1, 1]。
第 2 大的完美二叉子树的大小是 3。
*/
func kthLargestPerfectSubtree(root *TreeNode, k int) int {

	var dfs func(node *TreeNode) (bool, int)

	var perfect []int

	dfs = func(node *TreeNode) (ans bool, size int) {
		if node == nil {
			return true, 0
		}

		left, s1 := dfs(node.Left)
		right, s2 := dfs(node.Right)
		if left && right && s1 == s2 {
			perfect = append(perfect, s1+s2+1)
			return true, s1 + s2 + 1
		}
		return false, 0
	}
	dfs(root)
	sort.Slice(perfect, func(i, j int) bool {
		return perfect[i] > perfect[j]
	})
	if len(perfect) < k {
		return -1
	}
	return perfect[k-1]
}
