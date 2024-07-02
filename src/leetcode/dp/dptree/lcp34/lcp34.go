package main

/**
小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。小扣出于美观考虑，
希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，求所有染成蓝色的结点价值总和最大是多少？

思路：

1.考虑当前节点染不染色
2.如果当前节点染色然后枚举子节点可以染色的最大值，
3.当前节点不染色贼子节点可以染色最大数量为k

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {

	//1,2,3,4,5,6,
	var dfs func(*TreeNode, int) int

	memo := make([]map[*TreeNode]int, k+1)
	for i := range memo {
		memo[i] = make(map[*TreeNode]int)
	}
	dfs = func(root *TreeNode, t int) int {
		if root == nil {
			return 0
		}
		if _, ok := memo[t][root]; ok {
			return memo[t][root]
		}
		res := dfs(root.Left, k) + dfs(root.Right, k)
		for i := 0; i < t; i++ {
			res = max(res, dfs(root.Left, i)+dfs(root.Right, t-1-i)+root.Val)
		}
		memo[t][root] = res
		return res
	}
	return dfs(root, k)
}
