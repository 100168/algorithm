package profit

/*给定一个整数数组，其中第i个元素代表了第i天的股票价格 。
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func maxProfit5(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	//f[i]表示第i天结束之后的状态
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f0 := -prices[0]
	f1 := 0
	f2 := 0
	for i := 1; i < n; i++ {
		//可以保持前一天状态或者非冷冻期买入当前股票状态（取2者最大值）
		new0 := max(f0, f2-prices[i])
		//前一天卖出股票
		new1 := f0 + prices[i]
		new2 := max(f1, f2)

		f0, f1, f2 = new0, new1, new2
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
