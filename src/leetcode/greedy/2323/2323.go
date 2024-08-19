package main

import (
	"sort"
)

/**
给你两个 下标从 0 开始 的整数数组 jobs 和 相等 长度的 workers ，
其中 jobs[i]是完成第 i 个工作所需的时间，workers[j] 是第 j 个工人每天可以工作的时间。

每项工作都应该 正好 分配给一个工人，这样每个工人就 只能 完成一项工作。

返回分配后完成所有作业所需的最少天数。



示例 1:

输入: jobs = [5,2,4], workers = [1,7,5]
输出: 2
解释:
- 把第 2 个工人分配到第 0 个工作。他们花了 1 天时间完成这项工作。
- 把第 0 个工人分配到第 1 个工作。他们花了 2 天时间完成这项工作。
- 把第 1 个工人分配到第 2 个工作。他们花了 1 天时间完成这项工作。
所有工作完成需要 2 天，因此返回 2。
可以证明 2 天是最少需要的天数。

*/

func minimumTime(jobs []int, workers []int) int {
	sort.Ints(jobs)
	sort.Ints(workers)

	ans := 0
	for i, v := range jobs {

		ans = max(ans, (v-1)/workers[i]+1)
	}
	return ans
}
