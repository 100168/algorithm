package main

import "fmt"

/*
*
爱丽丝参与一个大致基于纸牌游戏 “21点” 规则的游戏，描述如下：

爱丽丝以 0 分开始，并在她的得分少于 k 分时抽取数字。 抽取时，她从 [1, maxPts] 的范围中随机获得一个整数作为分数进行累计，
其中 maxPts 是一个整数。 每次抽取都是独立的，其结果具有相同的概率。

当爱丽丝获得 k 分 或更多分 时，她就停止抽取数字。

爱丽丝的分数不超过 n 的概率是多少？

与实际答案误差不超过 10-5 的答案将被视为正确答案。

示例 1：

输入：n = 10, k = 1, maxPts = 10
输出：1.00000
解释：爱丽丝得到一张牌，然后停止。
示例 2：

输入：n = 6, k = 1, maxPts = 10
输出：0.60000
解释：爱丽丝得到一张牌，然后停止。 在 10 种可能性中的 6 种情况下，她的得分不超过 6 分。
*/
func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}
	for i := k - 1; i >= 0; i-- {
		for j := 1; j <= maxPts; j++ {
			dp[i] += dp[i+j] / float64(maxPts)
		}
	}
	fmt.Println(dp)

	return dp[0]
}

func new21Game2(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, k+maxPts)
	s := make([]float64, k+maxPts+1)
	//小于n的都为1且大于等于k的都为1
	for i := min(n, k+maxPts-1); i >= k; i-- {
		dp[i] = 1.0
		s[i] = s[i+1] + 1.0
	}
	//从后往前计算
	for i := k - 1; i >= 0; i-- {
		dp[i] = (s[i+1] - s[i+maxPts+1]) / float64(maxPts)
		s[i] = s[i+1] + dp[i]
	}
	return dp[0]
}

// 太难了。选择放弃
func new21Game3(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}
	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}

	dp[k-1] = 1.0 * float64(min(n-k+1, maxPts)) / float64(maxPts)
	for i := k - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+maxPts+1]-dp[i+1])/float64(maxPts)
	}
	return dp[0]
}

func main() {
	fmt.Println(new21Game2(5, 4, 4))
	fmt.Println(new21Game(5, 4, 4))
}
