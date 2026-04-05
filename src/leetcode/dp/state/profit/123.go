package profit

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][0] = -prices[0]
	//dp[0][1][0] = -prices[0]

	for i := 1; i < n; i++ {
		for k := 0; k < 2; k++ {

			if k > 0 {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k-1][1]-prices[i])

			} else {
				dp[i][0][0] = max(dp[i-1][0][0], -prices[i])
			}
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0]+prices[i])
		}
	}

	return dp[n-1][1][1]
}

func maxProfit3_(prices []int) int {
	n := len(prices)
	buy1 := -prices[0]
	buy2 := -prices[0]
	sell1 := 0
	sell2 := 0

	//dp[0][1][0] = -prices[0]

	for i := 1; i < n; i++ {
		newBuy1 := max(buy1, -prices[i])
		newBuy2 := max(buy2, sell1-prices[i])
		newSell1 := max(sell1, buy1+prices[i])
		newSell2 := max(sell2, buy2+prices[i])

		buy1, buy2, sell1, sell2 = newBuy1, newBuy2, newSell1, newSell2
	}

	return sell2
}
