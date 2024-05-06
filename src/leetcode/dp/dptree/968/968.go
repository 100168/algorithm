package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，我们在树的节点上安装摄像头。

节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
计算监控树的所有节点所需的最小摄像头数量。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int, int)

	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return math.MaxInt / 2, 0, 0
		}
		lChoose, lByFa, lByChild := dfs(node.Left)
		rChoose, rByFa, rByChild := dfs(node.Right)
		//当前节点装监控 == 子节点装或不装最小值相加
		choose := min(lChoose, lByFa) + min(rChoose, rByFa) + 1
		//当前节点被父节点监控==子节点监控或者子节点被自己儿子监控相加
		byFa := min(lChoose, lByChild) + min(rChoose, rByChild)
		byChild := min(lChoose+min(rChoose, rByChild), rChoose+min(lChoose, lByChild))
		return choose, byFa, byChild
	}
	choose, _, byChildren := dfs(root)
	return min(choose, byChildren)
}
