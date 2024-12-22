package main

/*
*
给你两个长度相同的字符串 s 和 t ，以及两个整数数组 nextCost 和 previousCost 。

一次操作中，你可以选择 s 中的一个下标 i ，执行以下操作 之一 ：

将 s[i] 切换为字母表中的下一个字母，如果 s[i] == 'z' ，切换后得到 'a' 。操作的代价为 nextCost[j] ，其中 j 表示 s[i] 在字母表中的下标。
将 s[i] 切换为字母表中的上一个字母，如果 s[i] == 'a' ，切换后得到 'z' 。操作的代价为 previousCost[j] ，其中 j 是 s[i] 在字母表中的下标。
切换距离 指的是将字符串 s 变为字符串 t 的 最少 操作代价总和。

请你返回从 s 到 t 的 切换距离 。

示例 1：

输入：s = "abab", t = "baba", nextCost = [100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], previousCost = [1,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]

输出：2

解释：

选择下标 i = 0 并将 s[0] 向前切换 25 次，总代价为 1 。
选择下标 i = 1 并将 s[1] 向后切换 25 次，总代价为 0 。
选择下标 i = 2 并将 s[2] 向前切换 25 次，总代价为 1 。
选择下标 i = 3 并将 s[3] 向后切换 25 次，总代价为 0 。

注意：竞赛中，请勿复制题面内容，以免影响您的竞赛成绩真实性。
*/
func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {

	ans := 0
	for i := range s {

		tt := s[i] - 'a'
		c := 0

		for tt != t[i]-'a' {
			tt++
			tt %= 26
			c += nextCost[tt]
		}
		cc := c
		c = 0

		tt = s[i] - 'a'

		for tt != t[i]-'a' {
			tt--
			tt = (tt + 26) % 26
			c += previousCost[tt]
		}
		ans += min(cc, c)

	}
	return int64(ans)

}
