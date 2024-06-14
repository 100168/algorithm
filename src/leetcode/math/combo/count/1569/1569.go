package main

import "fmt"

/*
*

给你一个数组 nums 表示 1 到 n 的一个排列。我们按照元素在 nums 中的顺序依次插入一个初始为空的二叉搜索树（BST）。
请你统计将 nums 重新排序后，统计满足如下条件的方案数：重排后得到的二叉搜索树与 nums 原本数字顺序得到的二叉搜索树相同。
比方说，给你 nums = [2,1,3]，我们得到一棵 2 为根，1 为左孩子，3 为右孩子的树。
数组 [2,3,1] 也能得到相同的 BST，但 [3,2,1] 会得到一棵不同的 BST 。

请你返回重排 nums 后，与原数组 nums 得到相同二叉搜索树的方案数。

由于答案可能会很大，请将结果对 10^9 + 7 取余数。

输入：nums = [3,4,5,1,2]
输出：5
解释：下面 5 个数组会得到相同的 BST：
[3,1,2,4,5]
[3,1,4,2,5]
[3,1,4,5,2]
[3,4,1,2,5]
[3,4,1,5,2]
*/
type node struct {
	value int
	left  *node
	right *node
}

func (t *node) insert(v int) {
	cur := t
	pre := t
	for cur != nil {
		pre = cur
		if cur.value > v {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if pre.value > v {
		pre.left = &node{v, nil, nil}
	} else {
		pre.right = &node{v, nil, nil}
	}
}
func numOfWays(nums []int) int {

	n := len(nums)
	mod := int(1e9 + 7)
	comb := make([][]int, n+1)
	for i := range comb {
		comb[i] = make([]int, n+1)
	}
	comb[0][0] = 1
	for i := 1; i <= n; i++ {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % mod
		}
	}
	root := &node{nums[0], nil, nil}

	for _, v := range nums[1:] {
		root.insert(v)
	}

	var dfs func(node *node) (int, int)
	dfs = func(node *node) (int, int) {

		if node == nil {
			return 0, 1
		}

		left, leftWays := dfs(node.left)
		right, rightWays := dfs(node.right)
		cur := left + right + 1
		curWays := comb[left+right][left] * leftWays % mod * rightWays % mod
		return cur, curWays
	}

	_, ways := dfs(root)
	return (ways - 1 + mod) % mod

}

func main() {

	fmt.Println(numOfWays([]int{3, 4, 5, 1, 2}))
}
