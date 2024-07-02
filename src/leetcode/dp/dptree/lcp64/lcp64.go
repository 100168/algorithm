package main

/***
力扣嘉年华」的中心广场放置了一个巨型的二叉树形状的装饰树。
每个节点上均有一盏灯和三个开关。节点值为 0 表示灯处于「关闭」状态，节点值为 1 表示灯处于「开启」状态。每个节点上的三个开关各自功能如下：
开关 1：切换当前节点的灯的状态；
开关 2：切换 以当前节点为根 的子树中，所有节点上的灯的状态；
开关 3：切换 当前节点及其左右子节点（若存在的话） 上的灯的状态；
给定该装饰的初始状态 root，请返回最少需要操作多少次开关，可以关闭所有节点的灯。

思路:
1.change
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closeLampInTree(root *TreeNode) (ans int) {
	type tuple struct {
		node             *TreeNode
		switch2, switch3 bool
	}
	dp := map[tuple]int{} // 记忆化搜索
	var f func(*TreeNode, bool, bool) int
	f = func(node *TreeNode, switch2, switch3 bool) int {
		if node == nil {
			return 0
		}
		p := tuple{node, switch2, switch3}
		if res, ok := dp[p]; ok {
			return res
		}
		//当前是1   switch2 和 switch3 可以抵消
		//当前是0   switch2 和 switch3 无法抵消
		if node.Val == 1 == (switch2 == switch3) { // 当前节点为开灯
			res1 := f(node.Left, switch2, false) + f(node.Right, switch2, false) + 1
			res2 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 1
			res3 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 1
			r123 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 3
			dp[p] = min(res1, res2, res3, r123)
		} else { // 当前节点为关灯
			res0 := f(node.Left, switch2, false) + f(node.Right, switch2, false)
			res12 := f(node.Left, !switch2, false) + f(node.Right, !switch2, false) + 2
			res13 := f(node.Left, switch2, true) + f(node.Right, switch2, true) + 2
			res23 := f(node.Left, !switch2, true) + f(node.Right, !switch2, true) + 2
			dp[p] = min(res0, res12, res13, res23)
		}
		return dp[p]
	}
	return f(root, false, false)
}
