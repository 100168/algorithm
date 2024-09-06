package main

/*
*
给你一个 正整数 数组 nums。

Alice 和 Bob 正在玩游戏。在游戏中，Alice 可以从 nums 中选择所有个位数 或 所有两位数，
剩余的数字归 Bob 所有。如果 Alice 所选数字之和 严格大于 Bob 的数字之和，则 Alice 获胜。

如果 Alice 能赢得这场游戏，返回 true；否则，返回 false。

示例 1：

输入：nums = [1,2,3,4,10]

输出：false

解释：

Alice 不管选个位数还是两位数都无法赢得比赛。

示例 2：

输入：nums = [1,2,3,4,5,14]

输出：true
*/
func canAliceWin(nums []int) bool {

	s1 := 0

	s2 := 0

	s := 0

	for _, v := range nums {
		if v < 100 && v > 9 {
			s2 += v
		} else if v < 10 {
			s1 += v
		}
		s += v
	}
	return max(s1, s2) > s-max(s1, s2)

}
