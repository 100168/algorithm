package main

import "math"

/*
*
给你一个下标从 0 开始的二进制数组 nums，其长度为 n ；另给你一个 正整数 k 以及一个 非负整数 maxChanges 。

Alice 在玩一个游戏，游戏的目标是让 Alice 使用 最少 数量的 行动 次数从 nums 中拾起 k 个 1 。
游戏开始时，Alice 可以选择数组 [0, n - 1] 范围内的任何索引 aliceIndex 站立。
如果 nums[aliceIndex] == 1 ，Alice 会拾起一个 1 ，并且 nums[aliceIndex] 变成0（这 不算 作一次行动）。
之后，Alice 可以执行 任意数量 的 行动（包括零次），在每次行动中 Alice 必须 恰好 执行以下动作之一：

选择任意一个下标 j != aliceIndex 且满足 nums[j] == 0 ，然后将 nums[j] 设置为 1 。这个动作最多可以执行 maxChanges 次。
选择任意两个相邻的下标 x 和 y（|x - y| == 1）且满足 nums[x] == 1, nums[y] == 0 ，然后交换它们的值（将 nums[y] = 1 和 nums[x] = 0）。
如果 y == aliceIndex，在这次行动后 Alice 拾起一个 1 ，并且 nums[y] 变成 0 。
返回 Alice 拾起 恰好 k 个 1 所需的 最少 行动次数。

思路: 中位数贪心+二分

//todo 先欠着，回去研究
*/
func minimumMoves(nums []int, k, maxChanges int) int64 {
	var pos []int
	c := 0 // nums 中连续的 1	 长度
	for i, x := range nums {
		if x == 0 {
			continue
		}
		pos = append(pos, i) // 记录 1 的位置
		c = max(c, 1)
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				c = 3 // 有 3 个连续的 1
			} else {
				c = max(c, 2) // 有 2 个连续的 1
			}
		}
	}

	c = min(c, k)
	if maxChanges >= k-c {
		// 其余 k-c 个 1 可以全部用两次操作得到
		return int64(max(c-1, 0) + (k-c)*2)
	}

	n := len(pos)
	sum := make([]int, n+1)
	for i, x := range pos {
		sum[i+1] = sum[i] + x
	}

	ans := math.MaxInt
	// 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
	size := k - maxChanges
	for right := size; right <= n; right++ {
		// s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 pos[(left+right)/2] 的距离之和
		left := right - size
		i := left + size/2
		s1 := pos[i]*(i-left) - (sum[i] - sum[left])
		s2 := sum[right] - sum[i] - pos[i]*(right-i)
		ans = min(ans, s1+s2)
	}
	return int64(ans + maxChanges*2)
}
