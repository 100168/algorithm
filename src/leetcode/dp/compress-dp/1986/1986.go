package main

import "math"

/*
*
你被安排了 n 个任务。任务需要花费的时间用长度为 n 的整数数组 tasks 表示，第 i 个任务需要花费 tasks[i] 小时完成。
一个 工作时间段 中，你可以 至多 连续工作 sessionTime 个小时，然后休息一会儿。
你需要按照如下条件完成给定任务：
如果你在某一个时间段开始一个任务，你需要在 同一个 时间段完成它。
完成一个任务后，你可以 立马 开始一个新的任务。
你可以按 任意顺序 完成任务。
给你 tasks 和 sessionTime ，请你按照上述要求，返回完成所有任务所需要的 最少 数目的 工作时间段 。
测试数据保证 sessionTime 大于等于 tasks[i] 中的 最大值 。
*/

func minSessions(tasks []int, sessionTime int) int {

	n := len(tasks)
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1
	}

	//先预处理，要不然会卡常
	sum := make([]int, 1<<n)
	for i, t := range tasks {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + t
		}
	}
	var dfs func(int) int
	dfs = func(mask int) int {

		if mask == 0 {
			return 0
		}

		if memo[mask] != -1 {
			return memo[mask]
		}

		ans := math.MaxInt

		for next := mask; next > 0; next = (next - 1) & mask {
			cur := 0
			for i := 0; i < n; i++ {
				if 1<<i&next != 0 {
					cur += tasks[i]
				}
			}
			if sum[next] <= sessionTime {
				ans = min(ans, dfs(mask^next)+1)
			}
		}
		memo[mask] = ans
		return ans
	}
	return dfs(1<<n - 1)
}

func minSessions2(tasks []int, sessionTime int) (ans int) {
	n := len(tasks)
	m := 1 << n
	// 预处理所有子集的子集和，复杂度 O(1+2+4+...+2^(n-1)) = O(2^n)
	sum := make([]int, m)
	for i, t := range tasks {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + t
		}
	}
	f := make([]int, m)
	for i := range f {
		f[i] = n
	}
	f[0] = 0
	for s := range f {
		// 枚举 s 的所有子集 sub，若 sub 耗时不超过 sessionTime，则将 f[s^sub]+1 转移到 f[s] 上
		for sub := s; sub > 0; sub = (sub - 1) & s {
			if sum[sub] <= sessionTime {
				f[s] = min(f[s], f[s^sub]+1)
			}
		}
	}
	return f[m-1]
}

func main() {
	println(minSessions([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 10))
	println(minSessions2([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 10))
}
