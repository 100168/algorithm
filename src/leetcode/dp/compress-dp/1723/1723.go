package main

/*
给你一个整数数组 jobs ，其中 jobs[i] 是完成第 i 项工作要花费的时间。

请你将这些工作分配给 k 位工人。所有工作都应该分配给工人，且每项工作只能分配给一位工人。
工人的 工作时间 是完成分配给他们的所有工作花费时间的总和。请你设计一套最佳的工作分配方案，使工人的 最大工作时间 得以 最小化 。

返回分配方案中尽可能 最小 的 最大工作时间 。
*/
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)

	sum := make([]int, 1<<n)
	for i, v := range jobs {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + v
		}
	}

	dp := make([]int, 1<<n)

	for i := range dp {
		dp[i] = sum[i]
	}

	for i := 1; i < k; i++ {
		for s := 1<<n - 1; s >= 0; s-- {
			for sub := s; sub > 0; sub = (sub - 1) & s {
				dp[s] = min(dp[s], max(dp[s^sub], sum[sub]))
			}
		}
	}

	return dp[1<<n-1]
}

func main() {
	println(minimumTimeRequired([]int{1, 2, 4, 7, 8}, 2))
}
