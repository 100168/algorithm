package main

import (
	"slices"
	"sort"
)

/*
*
给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。

商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。

返回礼盒的 最大 甜蜜度。

示例 1：

输入：price = [13,5,1,8,21,2], k = 3
输出：8
解释：选出价格分别为 [13,5,21] 的三类糖果。
礼盒的甜蜜度为 min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8 。
可以证明能够取得的最大甜蜜度就是 8 。
*/
func maximumTastiness(price []int, k int) int {

	l, r := 1, slices.Max(price)

	sort.Ints(price)
	for l <= r {
		m := (r + l) / 2

		if check(m, k, price) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

func check(t, k int, price []int) bool {

	cnt := 1
	pre := price[0]
	for i := 1; i < len(price); i++ {
		if price[i]-pre >= t {
			cnt++
			pre = price[i]
		}
	}
	return cnt >= k
}
