package profit

/*
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxProfit1(prices []int) int {

	n := len(prices)

	buy := -prices[0]
	sell := 0

	for i := 1; i < n; i++ {
		newSell := max(buy+prices[i], sell)

		newBuy := max(-prices[i], buy)

		buy, sell = newBuy, newSell
	}

	return sell
}
