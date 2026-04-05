package main

// 二叉树中，如果一个节点是其父节点的唯一子节点，则称这样的节点为 “独生节点” 。二叉树的根节点不会是独生节点，因为它没有父节点。
//
// 给定一棵二叉树的根节点root ，返回树中 所有的独生节点的值所构成的数组 。数组的顺序 不限
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-all-the-lonely-nodes
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func getLonelyNodes(root *TreeNode) []int {

	var ans []int
	ans = append(ans, dfs(root)...)
	return ans
}

func dfs(node *TreeNode) (ans []int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right != nil {
		ans = append(ans, node.Right.Val)
	}
	if node.Left != nil && node.Right == nil {
		ans = append(ans, node.Left.Val)
	}
	if node.Left == nil && node.Right == nil {
		return
	}

	ans = append(ans, dfs(node.Left)...)
	ans = append(ans, dfs(node.Right)...)
	return ans
}
