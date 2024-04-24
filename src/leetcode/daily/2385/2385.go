package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {

	g := make(map[int][]int)
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			g[node.Val] = append(g[node.Val], node.Left.Val)
			g[node.Left.Val] = append(g[node.Left.Val], node.Val)
			dfs(node.Left)
		}
		if node.Right != nil {
			g[node.Val] = append(g[node.Val], node.Right.Val)
			g[node.Right.Val] = append(g[node.Right.Val], node.Val)
			dfs(node.Right)
		}
	}

	var dfs2 func(int, int) int
	dfs2 = func(cur, p int) int {
		ans := 0
		for _, v := range g[cur] {
			if v != p {
				ans = max(ans, dfs2(v, cur)+1)
			}
		}
		return ans
	}

	dfs(root)
	return dfs2(start, -1)
}
