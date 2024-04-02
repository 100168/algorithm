package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	var dfs func(int, int) []*TreeNode

	cache := make([][][]*TreeNode, n+1)
	for i := range cache {
		cache[i] = make([][]*TreeNode, n+1)
	}

	dfs = func(left, right int) []*TreeNode {

		cur := make([]*TreeNode, 0)
		if left == right {
			cur = append(cur, &TreeNode{Val: left})
			return cur
		}
		if left > right {
			cur = append(cur, nil)
			return cur
		}
		if cache[left][right] != nil {
			return cache[left][right]
		}

		for root := left; root <= right; root++ {
			leftVal := dfs(left, root-1)
			rightVal := dfs(root+1, right)
			for _, l := range leftVal {
				for _, r := range rightVal {
					rootVal := &TreeNode{root, l, r}
					cur = append(cur, rootVal)
				}
			}
		}

		cache[left][right] = cur
		return cur

	}
	nodes := dfs(1, n)
	return nodes
}
func main() {
	println(generateTrees(1))
	println(generateTrees(2))
}
