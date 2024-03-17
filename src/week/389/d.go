package main

/*
*
给你一个下标从 0 开始的二进制数组 nums，其长度为 n ；另给你一个 正整数 k 以及一个 非负整数 maxChanges 。

灵茶山艾府在玩一个游戏，游戏的目标是让灵茶山艾府使用 最少 数量的 行动 次数从 nums 中拾起 k 个 1 。游戏开始时，灵茶山艾府可以选择数组 [0, n - 1] 范围内的任何索引index 站立。如果 nums[index] == 1 ，灵茶山艾府就会拾起一个 1 ，并且 nums[index] 变成0（这 不算 作一次行动）。之后，灵茶山艾府可以执行 任意数量 的 行动（包括零次），在每次行动中灵茶山艾府必须 恰好 执行以下动作之一：

选择任意一个下标 j != index 且满足 nums[j] == 0 ，然后将 nums[j] 设置为 1 。这个动作最多可以执行 maxChanges 次。
选择任意两个相邻的下标 x 和 y（|x - y| == 1）且满足 nums[x] == 1, nums[y] == 0 ，然后交换它们的值（将 nums[y] = 1 和 nums[x] = 0）。如果 y == index，在这次行动后灵茶山艾府拾起一个 1 ，并且 nums[y] 变成 0 。
返回灵茶山艾府拾起 恰好 k 个 1 所需的 最少 行动次数。
*/
func minimumMoves(nums []int, k int, maxChanges int) int64 {

	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int,int,int,int,bool) int
	dfs = func(index int, x, y ,j int, ops int,isOne bool) int {
		if isOne {
			if j== {
				
			}
		}
	}
}
