package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#,"
		}

		cur := strconv.Itoa(node.Val) + ","
		cur += dfs(node.Left)
		cur += dfs(node.Right)
		return cur
	}

	return pattern(dfs(root), dfs(subRoot))

}

// z函数
func pattern(s1, s2 string) bool {

	s1 = s1[:len(s1)-1]
	ss := s2 + s1
	s2 = s2[:len(s2)-1]
	m := len(strings.Split(s2, ","))
	sss := strings.Split(ss, ",")
	n := len(sss)
	z := make([]int, n)
	for i, c, r := 1, 1, 1; i < n; i++ {
		l := 0

		if i < r {
			l = min(r-i, z[i-c])
		}

		for i+l < n && sss[l] == sss[i+l] {
			l++
		}

		if l >= m {
			return true
		}
		z[i] = l
	}
	return false

}
