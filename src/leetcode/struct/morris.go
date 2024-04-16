package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func morris(node *TreeNode) {

	if node == nil {
		return
	}
	cur := node
	var mostRight *TreeNode
	for cur != nil {
		mostRight = cur.Left

		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right != cur {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}
		cur = cur.Right
	}
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	var mostRight *TreeNode

	var ans []int
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				ans = append(ans, cur.Val)
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			ans = append(ans, cur.Val)
		}

		cur = cur.Right
	}
	return ans
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	var mostRight *TreeNode

	var ans []int
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {

				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				ans = append(ans, cur.Val)
				mostRight.Right = nil
			}
		} else {
			ans = append(ans, cur.Val)
		}

		cur = cur.Right
	}
	return ans
}
