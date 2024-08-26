package main

import (
	"slices"
)

/*
*
给你一个长度为 n 的二维整数数组 items 和一个整数 k 。

items[i] = [profiti, categoryi]，其中 profiti 和 categoryi 分别表示第 i 个项目的利润和类别。

现定义 items 的 子序列 的 优雅度 可以用 total_profit + distinct_categories2 计算，
其中 total_profit 是子序列中所有项目的利润总和，distinct_categories 是所选子序列所含的所有类别中不同类别的数量。

你的任务是从 items 所有长度为 k 的子序列中，找出 最大优雅度 。

用整数形式表示并返回 items 中所有长度恰好为 k 的子序列的最大优雅度。

注意：数组的子序列是经由原数组删除一些元素（可能不删除）而产生的新数组，且删除不改变其余元素相对顺序。

思路：
1.先按种类排序
2.先选前k个元素。
3.然后再考虑要不要选当前元素，如果当前元素与之前元素种类一样，则跳过，否则选单前元素，并且替换之前选择的栈底元素

2.反悔贪心
*/
func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	var duplicate []int // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit // 累加前 k 个项目的利润
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !vis[category] { // 之前没有的类别
			vis[category] = true                                // len(vis) 变大
			totalProfit += profit - duplicate[len(duplicate)-1] // 选一个重复类别中的最小利润替换
			duplicate = duplicate[:len(duplicate)-1]
		} // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}
