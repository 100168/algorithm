package main

/*
*给定二叉树的根 root，具有以下属性:

叶节点 的值为 0 或 1，分别表示 false 和 true。
非叶节点的值为 2、3、4、5，分别表示布尔运算 OR, AND, XOR, NOT。
您还将得到一个布尔型 result，这是 root 节点的期望 评价 结果。

对节点的评价计算如下:

如果节点是叶节点，则评价是节点的 值，即 true 或 false.
否则, 将其值的布尔运算应用于子节点的 评价，该节点的 评价 即为布尔运算后的结果。
在一个操作中，您可以 翻转 一个叶节点，这将导致一个 false 节点变为 true 节点，一个 true 节点变为 false 节点。

返回需要执行的最小操作数，以使 root 的评价得到 result。可以证明，总有办法达到 result。

叶节点 是没有子节点的节点。

注意: NOT 节点只有左孩子或只有右孩子，但其他非叶节点同时拥有左孩子和右孩子。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumFlips(root *TreeNode, result bool) int {

	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		if node.Val == 1 {
			return 1, 0
		}
		if node.Val == 0 {
			return 0, 1
		}
		lFalse, lTrue := dfs(node.Left)
		rFalse, rTrue := dfs(node.Left)

		if node.Val == 5 {
			return lFalse + rFalse, lTrue + rTrue
		}
	}
	return 1
}

func main() {

}
