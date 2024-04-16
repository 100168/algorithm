package main

/*
给定一个二叉树（具有根结点root），一个目标结点target，和一个整数值 K 。

返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := make(map[int]*TreeNode)
	current := root

	var stack []*TreeNode

	var ans []int
	stack = append(stack, current)
	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		if pop == target {
			break
		}
		stack = stack[:len(stack)-1]
		if pop.Left != nil {
			stack = append(stack, pop.Left)
			parent[pop.Left.Val] = pop
		}
		if pop.Right != nil {
			stack = append(stack, pop.Right)
			parent[pop.Right.Val] = pop
		}
	}

	//用来表示已经遍历过的结点
	path := make(map[int]bool)
	ans = append(ans, getK(target, k, path)...)
	current = target
	path[target.Val] = true
	for parent[current.Val] != nil {
		k--
		current = parent[current.Val]
		ans = append(ans, getK(current, k, path)...)
	}

	return ans
}

func getK(root *TreeNode, k int, path map[int]bool) (ans []int) {
	if path[root.Val] {
		return
	}
	path[root.Val] = true
	if k == 0 {
		ans = append(ans, root.Val)
		return ans
	}
	if root.Left != nil {
		ans = append(ans, getK(root.Left, k-1, path)...)
	}
	if root.Right != nil {
		ans = append(ans, getK(root.Right, k-1, path)...)
	}

	return ans
}
