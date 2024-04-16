package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树[3,9,20,null,null,15,7],

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	return [][]int{{1}, {2}}
}

func processZig(flag bool, ans [][]int, node *TreeNode) {

}
func main() {

}
