package rob

/*在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := process(root)
	return max(ans[0], ans[1])
}

func process(node *TreeNode) []int {

	if node == nil {
		return []int{0, 0}
	}

	l, r := process(node.Left), process(node.Right)

	//偷当前=当前+不偷left+不偷right
	selected := node.Val + l[1] + r[1]
	//不偷当前=
	noSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, noSelected}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
