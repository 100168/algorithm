package main

type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {

	ans := 0

	var dfs func(*Node) int

	dfs = func(node *Node) int {
		if node == nil {
			return -1
		}

		cur := 0
		for _, v := range node.Children {
			l := dfs(v) + 1
			ans = max(ans, l+cur)
			cur = max(cur, l)
		}
		return cur
	}

	dfs(root)
	return ans
}
