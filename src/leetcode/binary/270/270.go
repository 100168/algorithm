package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	ans := 1<<31 - 1

	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(ans)-target) {
			ans = root.Val
		}
		if float64(root.Val) > target {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return ans

}
