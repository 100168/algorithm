package main

import "fmt"

/*数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func minCostClimbingStairs(cost []int) int {

	for _, i2 := range cost {
		fmt.Println(i2)
	}
	n := len(cost)
	a := 0
	b := 0
	for i := 2; i <= n; i++ {
		c := b
		if a+cost[i-2] < b+cost[i-1] {
			b = a + cost[i-2]
		} else {
			b = b + cost[i-1]
		}
		a = c
	}
	return b
}

// 到每一步都往前看一步到当前位置的最小消耗.
func minCost(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if dp[i-2]+cost[i-2] < dp[i-1]+cost[i-1] {
			dp[i] = dp[i-2] + cost[i-2]
		} else {
			dp[i] = dp[i-1] + cost[i-1]
		}
	}
	return dp[n]
}
