package main

import "sort"

/*
*
给你一个正整数 days，表示员工可工作的总天数（从第 1 天开始）。另给你一个二维数组 meetings，长度为 n，其中 meetings[i] = [start_i, end_i] 表示第 i 次会议的开始和结束天数（包含首尾）。

返回员工可工作且没有安排会议的天数。

注意：会议时间可能会有重叠。
*/
func countDays(days int, meetings [][]int) int {

	ans := 0
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	pre := 1

	for _, m := range meetings {
		start, end := m[0], m[1]

		ans += max(0, start-pre-1)
		pre = max(pre, end)

	}
	return ans + max(pre, days-pre-1)
}
