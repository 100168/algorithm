package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {

	var dfs func(int, int) *TreeNode

	dfs = func(l, r int) *TreeNode {

		if l > r {
			return nil
		}
		t := new(TreeNode)
		t.Val = preorder[l]
		if l == r {

			return t
		}

		m := l

		for i := l + 1; i <= r; i++ {

			if preorder[i] > preorder[l] {
				m = i
				break
			}
		}

		if m != l {
			t.Left = dfs(l+1, m-1)
		} else {
			t.Left = dfs(l+1, r)
		}

		if m > l {
			t.Right = dfs(m, r)
		}

		return t

	}

	return dfs(0, len(preorder)-1)
}
func main() {
	//fmt.Println(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
	fmt.Println(bstFromPreorder([]int{4, 2}))
}
