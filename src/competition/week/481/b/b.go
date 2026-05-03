package main

/*
*
给你一个长度为 n 的字符串 s 和一个整数数组 cost，其中 cost[i] 表示删除字符串 s 中第 i 个字符的代价。

你可以从字符串 s 中删除任意数量的字符（也可以不删除），使得最终的字符串非空且由相同字符组成。

返回实现上述目标所需的最小总删除代价。

示例 1：

输入： s = "aabaac", cost = [1,2,3,4,1,10]

输出： 11

解释：

删除索引为 0、1、2、3 和 4 的字符后，字符串变为 "c"，它由相同的字符组成，总删除代价为：cost[0] + cost[1] + cost[2] + cost[3] + cost[4] = 1 + 2 + 3 + 4 + 1 = 11。

示例 2：

输入： s = "abc", cost = [10,5,8]

输出： 13

解释：

删除索引为 1 和 2 的字符后，字符串变为 "a"，它由相同的字符组成，总删除代价为：cost[1] + cost[2] = 5 + 8 = 13。

示例 3：

输入： s = "zzzzz", cost = [67,67,67,67,67]

输出： 0

解释：

字符串 s 中的所有字符都相同，因此不需要删除字符，删除代价为 0。

提示：

n == s.length == cost.length
1 <= n <= 10^5
1 <= cost[i] <= 10^9
s 仅由小写英文字母组成。
*/
func minCost(s string, cost []int) int64 {

	total := int64(0)
	cnt := make([]int64, 26)
	for i, c := range s {
		val := int64(cost[i])
		total += val
		cnt[c-'a'] += val
	}
	maxKeep := int64(0)
	for _, v := range cnt {
		if v > maxKeep {
			maxKeep = v
		}
	}
	return total - maxKeep

}
