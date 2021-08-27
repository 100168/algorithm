package main

/*给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
否则不为NULL 的节点将直接作为新二叉树的节点

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-binary-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var pointer1 *TreeNode
	var pointer2 *TreeNode

	pointer1 = root1
	pointer2 = root2
	return process(pointer1, pointer2)
}
func process(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var resPointer *TreeNode
	if root1 != nil && root2 != nil {
		root1.Val = root1.Val + root2.Val
		resPointer = root1
		resPointer.Left = process(root1.Left, root2.Left)
		resPointer.Right = process(root1.Right, root2.Right)
	} else if root1 != nil {
		resPointer = root1
	} else {
		resPointer = root2
	}
	return resPointer

}

func main() {

}
