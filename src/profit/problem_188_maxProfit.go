package profit

import "math"

/*给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func maxProfit4(k int, prices []int) int {
	n := len(prices)

	k = min(k, n/2)
	dp := make([][][]int, n)

	if n == 0 || k == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -prices[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			if j > 0 {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]-prices[i])
			} else {
				dp[i][0][0] = max(dp[i-1][0][0], -prices[i])
			}

			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]+prices[i])
		}
	}

	return dp[n-1][k-1][1]
}

func maxProfit(k int, prices []int) int {
	n := len(prices)

	k = min(k, n/2)
	buy := make([]int, k)
	sell := make([]int, k)

	if n == 0 || k == 0 {
		return 0
	}
	for i := 1; i < k; i++ {
		buy[i] = math.MinInt16
		sell[i] = math.MinInt16
	}
	buy[0] = -prices[0]
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < k; j++ {
			sell[j] = max(sell[j], buy[j]+prices[i])
			buy[j] = max(buy[j], sell[j-1]-prices[i])
		}
	}
	return max1(sell...)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max1(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
