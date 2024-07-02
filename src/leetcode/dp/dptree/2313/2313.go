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

	type pair struct {
		node   *TreeNode
		status bool
	}
	dp := make(map[pair]int)
	var dfs func(*TreeNode, bool) int
	dfs = func(node *TreeNode, status bool) int {
		if node == nil {
			return 0
		}
		p := pair{node, status}
		if _, ok := dp[p]; ok {
			return dp[p]
		}

		if node.Val == 1 {
			if status {
				dp[p] = 0
			} else {
				dp[p] = 1
			}
			return dp[p]
		}
		if node.Val == 0 {
			if status {
				dp[p] = 1
			} else {
				dp[p] = 0
			}
			return dp[p]
		}
		cur := 0
		/*
			非叶节点的值为 2、3、4、5，分别表示布尔运算 OR, AND, XOR, NOT。*/
		if status {
			switch node.Val {
			case 2:
				cur = min(dfs(node.Left, true), dfs(node.Right, true))
			case 3:
				cur = dfs(node.Left, true) + dfs(node.Right, true)
			case 4:
				cur = min(dfs(node.Left, false)+dfs(node.Right, true), dfs(node.Left, true)+dfs(node.Right, false))
			case 5:
				if node.Left != nil {
					cur = dfs(node.Left, false)
				} else {
					cur = dfs(node.Right, false)
				}
			}
		} else {
			switch node.Val {
			case 2:
				cur = min(dfs(node.Left, false) + dfs(node.Right, false))
			case 3:
				cur = min(dfs(node.Left, false), dfs(node.Right, false))
			case 4:
				cur = min(dfs(node.Left, false)+dfs(node.Right, false), dfs(node.Left, true)+dfs(node.Right, true))
			case 5:
				if node.Left != nil {
					cur = dfs(node.Left, true)
				} else {
					cur = dfs(node.Right, true)
				}
			}
		}
		dp[p] = cur
		return cur
	}
	return dfs(root, result)
}

func main() {

}
