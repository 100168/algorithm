package main

import (
	"math"
	"math/bits"
)

/**
Bob 被困在了一个地窖里，他需要破解 n 个锁才能逃出地窖，每一个锁都需要一定的 能量 才能打开。每一个锁需要的能量存放在一个数组 strength 里，其中 strength[i] 表示打开第 i 个锁需要的能量。

Bob 有一把剑，它具备以下的特征：

一开始剑的能量为 0 。
剑的能量增加因子 X 一开始的值为 1 。
每分钟，剑的能量都会增加当前的 X 值。
打开第 i 把锁，剑的能量需要到达 至少 strength[i] 。
打开一把锁以后，剑的能量会变回 0 ，X 的值会增加一个给定的值 K 。
你的任务是打开所有 n 把锁并逃出地窖，请你求出需要的 最少 分钟数。

请你返回 Bob 打开所有 n 把锁需要的 最少 时间。


*/

func findMinimumTime(strength []int, K int) int {

	n := len(strength)

	f := make([]int, 1<<n)

	for i := range f {
		f[i] = -1
	}

	var dfs func(int) int

	dfs = func(mask int) int {

		if mask == 1<<n-1 {
			return 0
		}

		c := bits.OnesCount(uint(mask))

		if f[mask] != -1 {
			return f[mask]
		}

		cur := math.MaxInt / 2

		for j := 0; j < n; j++ {

			if mask>>j&1 == 0 {
				cost := (strength[j]-1)/(c*K+1) + 1
				cur = min(cur, dfs(1<<j|mask)+cost)
			}
		}

		f[mask] = cur
		return cur
	}

	return dfs(0)

}
