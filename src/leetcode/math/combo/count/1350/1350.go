package main

import "fmt"

/*
*
给你 n 笔订单，每笔订单都需要快递服务。

计算所有有效的 取货 / 交付 可能的顺序，使 delivery(i) 总是在 pickup(i) 之后。

由于答案可能很大，请返回答案对 10^9 + 7 取余的结果。

输入：n = 2
输出：6
解释：所有可能的序列包括：
(P1,P2,D1,D2)，(P1,P2,D2,D1)，(P1,D1,P2,D2)，(P2,P1,D1,D2)，(P2,P1,D2,D1) 和 (P2,D2,P1,D1)。
(P1,D2,P2,D1) 是一个无效的序列，因为物品 2 的收件服务（P2）不应在物品 2 的配送服务（D2）之后。

思路：
1.排列组合
2.从左往右计算，考虑到当前位置 1.放一起2.不放一起
*/
func countOrders(n int) int {

	mod := int(1e9 + 7)
	ans := 1
	for i := 1; i <= n; i++ {
		pre := 2*(i-1) + 1
		ans = ans * (pre + (pre * (pre - 1) / 2)) % mod
	}
	return ans
}

func main() {
	fmt.Println(countOrders(1))
}
