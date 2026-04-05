package main

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func coinChange(coins []int, amount int) int {

	n := len(coins)
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = 1 << 30
		min := 1 << 30
		for j := 0; j < n; j++ {
			diff := i - coins[j]
			if diff >= 0 {
				if dp[diff]+1 != 0 && dp[diff]+1 < min {
					min = dp[diff]
				}
			}
		}
		dp[i] = min
	}
	if dp[amount] == 1<<30 {
		return -1
	} else {
		return dp[amount]
	}
}
func main() {

}
